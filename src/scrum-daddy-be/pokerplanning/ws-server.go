package pokerplanning

import (
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/common/results"
	"scrum-daddy-be/common/utils"
	"scrum-daddy-be/pokerplanning/domain"
	"strconv"
	"sync"
)

type RegisterMessage struct {
	Client    *Client
	PokerRoom domain.PokerRoom
}

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
	ID        int                   `json:"id"`
	Clients   map[uuid.UUID]*Client `json:"clients"`
	PokerRoom domain.PokerRoom
}

type Hub struct {
	rooms      map[int]*HubRoom
	broadcast  chan *HubMessage
	register   chan *Client
	unregister chan *Client
	mtx        sync.RWMutex
	db         *db.Database
	ctx        context.Context
}

func NewRoomHub(ctx context.Context, db *db.Database) *Hub {
	return &Hub{
		rooms:      make(map[int]*HubRoom),
		broadcast:  make(chan *HubMessage, 50),
		register:   make(chan *Client, 5),
		unregister: make(chan *Client, 5),
		mtx:        sync.RWMutex{},
		db:         db,
		ctx:        ctx,
	}
}

func (hub *Hub) roomExists(roomId int) bool {
	_, ok := hub.rooms[roomId]
	return ok
}

func (hub *Hub) registerRoom(roomId int, room domain.PokerRoom) {
	_, ok := hub.rooms[roomId]
	if ok {
		return
	}
	hub.rooms[roomId] = &HubRoom{
		ID:        roomId,
		Clients:   make(map[uuid.UUID]*Client),
		PokerRoom: room,
	}
}

func (hub *Hub) listen() {
	slog.Info("Ws listening for messages...")
	for {
		select {
		case client := <-hub.register:
			hub.handleRegisterEvent(client)
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
				slog.Info("message from err handler")
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

func (hub *Hub) handleRegisterEvent(client *Client) {
	slog.Info("Registering client", "client", client.ID, "room", client.RoomID)
	r, ok := hub.rooms[client.RoomID]
	if !ok {
		slog.Error("Room not registered in rooms map", "roomId", client.RoomID)
		return
	}
	if _, ok := r.Clients[client.ID]; !ok {
		slog.Debug("Adding client to room")
		r.Clients[client.ID] = client

		m := &HubMessage{
			Type:   Join,
			RoomID: client.RoomID,
			User:   client.Username,
		}
		hub.broadcast <- m
	} else {
		// TODO: Handle client already in room
		slog.Debug("Client already in room")
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
	if err != nil {
		slog.Error("Error parsing roomId", "err", err)
		_ = conn.Close()
		_ = api.WriteJSON(w, http.StatusBadRequest, results.NewErrorResult(
			http.StatusBadRequest,
			"Invalid room id",
			err.Error()),
		)
		return
	}

	if !hub.roomExists(roomId) {
		repo := NewPokerRoomRepository(hub.db)
		room, errResult := GetPokerRoomBySecondaryId(
			hub.ctx,
			repo,
			roomId,
		)
		if errResult != nil {
			slog.Error("Error getting room", "err", errResult)
			_ = conn.Close()
			_ = api.WriteJSON(w, http.StatusNotFound, errResult)
			return
		}
		hub.registerRoom(roomId, *room)
	}

	slog.Debug("New client connected", "roomId", roomId, "userId", userId)
	username := "mikica " + strconv.Itoa(rand.Int())
	client := &Client{
		conn:     conn,
		RoomID:   roomId,
		Username: username,
		ID:       userId,
		Message:  make(chan *HubMessage, 10),
	}
	hub.register <- client

	go client.readMessage(hub)
	go client.writeMessage()
}
