package pokerplanning

import (
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/errors"
	"scrum-daddy-be/pokerplanning/abstractions"
	"scrum-daddy-be/pokerplanning/types"
)

type CreateRoomRequest struct {
	Name    string    `json:"name"`
	OwnerId uuid.UUID `json:"ownerId"`
}

// HandlePostPokerRooms @Summary Create new poker room
// @Description Create new poker room
// @Tags rooms/v1
// @Accept  json
// @Produce  json
// @Param   requestBody  body      CreateRoomRequest  true  "Create Room Request"
// @Success 200 {object} api.CreateResponse
// @Failure 400 {object} errors.ErrorResult
// @Router /api/v1/rooms [post]
func (c Container) HandlePostPokerRooms(w http.ResponseWriter, r *http.Request) *errors.ErrorResult {
	createRoomReq := new(CreateRoomRequest)
	if err := json.NewDecoder(r.Body).Decode(createRoomReq); err != nil {
		return errors.NewErrorResult(
			http.StatusBadRequest,
			"Invalid request",
			err.Error())
	}
	log.Println("Create room request", *createRoomReq)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)
	roomId, err := CreateRoom(c.IPokerRoomRepository, createRoomReq)
	if err != nil {
		return errors.NewErrorResult(
			http.StatusBadRequest,
			"Cannot create room",
			err.Error(),
		)
	}
	response := api.CreateResponse{Id: roomId}
	_ = api.WriteJSON(w, http.StatusCreated, response)
	return nil
}

func CreateRoom(repo abstractions.IPokerRoomRepository, request *CreateRoomRequest) (uuid.UUID, error) {
	room := &types.PokerRoom{
		Name:    request.Name,
		OwnerId: uuid.New(),
	}
	id, err := repo.Save(room)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
