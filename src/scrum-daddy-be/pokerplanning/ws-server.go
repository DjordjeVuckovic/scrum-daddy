package pokerplanning

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"scrum-daddy-be/common/utils"
	"strconv"
	"sync"
)

func upgradeToWs() *websocket.Upgrader {
	return &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // TODO Not Allow connections from any origin
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}

type HubRoom struct {
	ID      int                   `json:"id"`
	Name    string                `json:"name"`
	Clients map[uuid.UUID]*Client `json:"clients"`
}

type Hub struct {
	rooms      map[int]*HubRoom
	broadcast  chan *HubMessage
	register   chan *Client
	unregister chan *Client
	mtx        sync.RWMutex
}

func NewRoomHub() *Hub {
	return &Hub{
		rooms:      make(map[int]*HubRoom),
		broadcast:  make(chan *HubMessage, 50),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mtx:        sync.RWMutex{},
	}
}

func (hub *Hub) listen() {
	slog.Info("Ws listening for messages...")
	r := HubRoom{
		ID:      1,
		Name:    "room 1",
		Clients: make(map[uuid.UUID]*Client),
	}
	hub.rooms[r.ID] = &r
	for {
		select {
		case client := <-hub.register:
			slog.Info("Registering client", "client", client.ID, "room", client.RoomID)
			if r, ok := hub.rooms[client.RoomID]; ok {

				if _, ok := r.Clients[client.ID]; !ok {
					slog.Info("Adding client to room")
					r.Clients[client.ID] = client
				} else {
					slog.Info("Client already in room")
				}
			}
		case client := <-hub.unregister:
			slog.Debug("A user has left the room", "roomId", client.RoomID)
			if _, ok := hub.rooms[client.RoomID]; ok {
				if _, ok := hub.rooms[client.RoomID].Clients[client.ID]; ok {
					if len(hub.rooms[client.RoomID].Clients) != 0 {
						hub.broadcast <- &HubMessage{
							Type:   Left,
							User:   client.Username,
							RoomID: client.RoomID,
						}
					}
				}
				err := client.conn.Close()
				if err != nil {
					slog.Error("Error closing connection: ", "error", err)
				}
				slog.Debug("Removing client from room", "client", client.ID, "room", client.RoomID)
				delete(hub.rooms[client.RoomID].Clients, client.ID)
				close(client.Message)
			} else {
				slog.Info("Not okk unreg")
			}
		case m, ok := <-hub.broadcast:
			if !ok {
				slog.Error("Broadcast channel closed", "closed channel")
			}
			slog.Info("Broadcasting message", "message", m)
			if _, ok := hub.rooms[m.RoomID]; ok {
				slog.Info("Broadcasting message to room", "room", m.RoomID, "message", m.Type)
				for _, cl := range hub.rooms[m.RoomID].Clients {
					slog.Info("Broadcasting message to client", "client", cl.Username, "message", m.Type)
					cl.Message <- m
				}
			}
		}
	}
}

func (hub *Hub) serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgradeToWs().Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Error upgrading connection: ", err)
		return
	}
	userIdStr := r.Header.Get("Authorization")
	userId, err := uuid.Parse(userIdStr)

	var roomIdStr = r.PathValue("roomId")
	roomId, err := utils.ParseToInt(roomIdStr)

	slog.Info("New client connected")
	username := "mikica " + strconv.Itoa(rand.Int())
	client := &Client{
		conn:     conn,
		RoomID:   roomId,
		Username: username,
		ID:       userId,
		Message:  make(chan *HubMessage, 10),
	}
	m := &HubMessage{
		Type:   Join,
		RoomID: roomId,
		User:   username,
	}
	hub.register <- client
	hub.broadcast <- m

	go client.readMessage(hub)
	go client.writeMessage()
}
