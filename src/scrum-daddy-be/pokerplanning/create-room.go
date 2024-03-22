package pokerplanning

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/results"
	"scrum-daddy-be/pokerplanning/abstractions"
	"scrum-daddy-be/pokerplanning/domain"
)

type CreateRoomRequest struct {
	Name         string    `json:"name"`
	OwnerId      uuid.UUID `json:"ownerId"`
	VotingSystem string    `json:"votingSystem"`
	IsAllReveal  bool      `json:"isAllReveal"`
	AutoReveal   bool      `json:"autoReveal"`
	ShowAverage  bool      `json:"showAverage"`
}

// HandlePostPokerRooms @Summary Create new poker room
// @Detail Create new poker room
// @Tags rooms
// @Accept  json
// @Produce  json
// @Param   requestBody  body      CreateRoomRequest  true  "Create Room Request"
// @Success 200 {object} api.CreateResponse
// @Failure 400 {object} results.ErrorResult
// @Router /api/v1/rooms [post]
func (c Container) HandlePostPokerRooms(w http.ResponseWriter, r *http.Request) *results.ErrorResult {
	createRoomReq := new(CreateRoomRequest)
	if err := json.NewDecoder(r.Body).Decode(createRoomReq); err != nil {
		return results.NewErrorResult(
			http.StatusBadRequest,
			"Invalid request",
			err.Error())
	}
	log.Println("Create room request", *createRoomReq)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	repo := NewPokerRoomRepository(c.Db)
	roomId, err := CreateRoom(repo, createRoomReq)

	if err != nil {
		var errResult *results.ErrorResult
		errors.As(err, &errResult)
		return errResult
	}
	response := api.CreateResponse{Id: roomId}
	_ = api.WriteJSON(w, http.StatusCreated, response)
	return nil
}

func CreateRoom(
	repo abstractions.IPokerRoomRepository,
	request *CreateRoomRequest) (uuid.UUID, error) {
	room, err := domain.NewPokerRoom(request.Name, request.OwnerId)
	if err != nil {
		return uuid.Nil, err
	}
	id, err := repo.Save(room)

	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
