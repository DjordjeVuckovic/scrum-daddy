package pokerplanning

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"log/slog"
	"net/http"
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

type RoomManager struct {
	clients      map[*Client]bool
	rooms        map[string]map[*Client]bool
	broadcast    chan Message
	register     chan *Client
	unregister   chan *Client
	clientsMutex sync.RWMutex
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		clients:      make(map[*Client]bool),
		rooms:        make(map[string]map[*Client]bool),
		broadcast:    make(chan Message),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		clientsMutex: sync.RWMutex{},
	}
}

func (manager *RoomManager) listen() {
	slog.Info("Ws listening for messages...")
	for {
		select {
		case client := <-manager.register:
			manager.clients[client] = true
			if _, ok := manager.rooms[client.roomId]; !ok {
				manager.rooms[client.roomId] = make(map[*Client]bool)
			}
			manager.rooms[client.roomId][client] = true
			// Notify clients in the same roomId
			log.Println("A new user has joined the roomId: ", client.roomId)
			manager.broadcast <- Message{Type: "join", User: "A new user has joined the roomId: " + client.roomId}
		case client := <-manager.unregister:
			slog.Debug("A user has left the roomId: ", "client", client.roomId)
			if _, ok := manager.clients[client]; ok {
				delete(manager.clients, client)
				delete(manager.rooms[client.roomId], client)
				err := client.conn.Close()
				if err != nil {
					slog.Error("Error closing connection: ", "error", err)
				}
			}
		case message, ok := <-manager.broadcast:
			if !ok {
				slog.Error("Broadcast channel closed")
			}
			slog.Debug("Broadcast message received: ", "message", message)
			for client := range manager.clients {
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
						delete(manager.clients, client)
					}
				}
			}
		}
	}
}

func (manager *RoomManager) serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgradeToWs().Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Error upgrading connection: ", "error", err)
		return
	}
	client := &Client{conn: conn, roomId: "exampleRoom"}
	manager.register <- client

	go func() {
		defer func() {
			manager.unregister <- client
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
			manager.broadcast <- msg
			slog.Debug("Message brod: ", "message", msg)
		}
	}()
}
