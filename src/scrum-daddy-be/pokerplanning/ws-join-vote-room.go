package pokerplanning

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"log/slog"
	"net/http"
	"scrum-daddy-be/common/logger"
	"sync"
)

type MessageType string

const (
	Join MessageType = "join"
	Vote MessageType = "vote"
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

type Client struct {
	conn   *websocket.Conn
	roomId string
	user   RoomUser
}

type RoomUser struct {
	Username string    `json:"username"`
	ID       uuid.UUID `json:"id"`
}

type Message struct {
	Type MessageType `json:"type"` // "join" or "vote"
	User string      `json:"user"`
	Vote string      `json:"vote,omitempty"` // Included for vote messages
}

type RoomHub struct {
	clients      map[*Client]bool
	rooms        map[string]map[*Client]bool
	broadcast    chan Message
	register     chan *Client
	removeClient chan *Client
	clientsMutex sync.RWMutex
}

func NewRoomManager() *RoomHub {
	return &RoomHub{
		clients:      make(map[*Client]bool),
		rooms:        make(map[string]map[*Client]bool),
		broadcast:    make(chan Message),
		register:     make(chan *Client),
		removeClient: make(chan *Client),
		clientsMutex: sync.RWMutex{},
	}
}

func (hub *RoomHub) listen() {
	slog.Info("Ws listening for messages...")
	for {
		select {
		case client := <-hub.register:
			hub.clients[client] = true
			if _, ok := hub.rooms[client.roomId]; !ok {
				hub.rooms[client.roomId] = make(map[*Client]bool)
			}
			hub.rooms[client.roomId][client] = true
			// Notify clients in the same roomId
			log.Println("A new user has joined the roomId: ", client.roomId)
			hub.broadcast <- Message{Type: Join, User: "A new user has joined the roomId: " + client.roomId}
		case client := <-hub.removeClient:
			slog.Debug("A user has left the roomId: ", "client", client.roomId)
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				delete(hub.rooms[client.roomId], client)
				err := client.conn.Close()
				if err != nil {
					slog.Error("Error closing connection: ", "error", err)
				}
			}
		case message, ok := <-hub.broadcast:
			if !ok {
				slog.Error("Broadcast channel closed")
			}
			slog.Debug("Broadcast message received: ", "message", message)
			for client := range hub.clients {
				err := client.conn.WriteJSON(message)
				log.Println("message: ", message, "client: ", client)
				if err != nil {
					slog.Error("Error writing message: ", "error", err)

					if websocket.IsUnexpectedCloseError(
						err,
						websocket.CloseGoingAway,
						websocket.CloseAbnormalClosure) {
						err := client.conn.Close()
						if err != nil {
							slog.Error("Error closing ws connection: ", "error", err)
						}
						delete(hub.clients, client)
					}
				}
			}
		}
	}
}

func (hub *RoomHub) serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgradeToWs().Upgrade(w, r, nil)
	if err != nil {
		logger.Error("Error upgrading connection: ", err)
		return
	}
	client := &Client{conn: conn, roomId: "exampleRoom"}
	hub.register <- client

	go func() {
		defer func() {
			hub.removeClient <- client
		}()
		for {
			var msg Message
			err := conn.ReadJSON(&msg)
			log.Println("msg read: ", msg)
			if err != nil {
				slog.Error("error while reading message", "error", err)
				break
			}
			slog.Debug("Message received: ", "message", msg)
			hub.broadcast <- msg
			slog.Debug("Message brod: ", "message", msg)
		}
	}()
}
