package pokerplanning

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log/slog"
)

type MessageType string

type Client struct {
	conn     *websocket.Conn
	RoomID   int       `json:"roomId"`
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Message  chan *HubMessage
}

const (
	Join MessageType = "join"
	Vote MessageType = "vote"
	Left MessageType = "left"
)

//type HubMessagePayload struct {
//	Type MessageType `json:"type"`
//}

type HubMessage struct {
	Type   MessageType `json:"type"` // "join" or "vote"
	User   string      `json:"user"`
	RoomID int         `json:"roomId"`
	Vote   int         `json:"vote,omitempty"`
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.unregister <- c
		_ = c.conn.Close()
	}()

	for {
		msg := new(HubMessage)
		err := c.conn.ReadJSON(msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure) {
				slog.Error("error while reading message", "err", err)
			}
			break
		}
		slog.Debug("HubMessage received: ", "message", msg)
		hub.broadcast <- msg
	}
}

func (c *Client) writeMessage() {
	defer func() {
		err := c.conn.Close()
		if err != nil {
			slog.Error("Error closing connection", "err", err)
		}
	}()
	for {
		slog.Debug("Writing message", "client", c.ID, "room", c.RoomID)
		message, ok := <-c.Message
		if !ok {
			slog.Error("Error reading message")
			return
		}

		_ = c.conn.WriteJSON(message)
	}
}
