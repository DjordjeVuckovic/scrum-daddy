package pokerplanning

import (
	"net/http"
	"scrum-daddy-be/common/results"
)

type JoinPokerRoomRequest struct {
	UserId string `json:"userId"`
	RoomId string `json:"roomId"`
}

func (c Container) HandleJoinPokerRoom(w http.ResponseWriter, r *http.Request) *results.ErrorResult {
	return nil
}
