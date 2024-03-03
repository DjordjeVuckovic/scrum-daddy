package rooms

import (
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/common/errors"
)

type Api struct {
	Db     *db.Database
	Server *api.Server
	IPokerRoomRepository
}

func AddPokerRooms(a Api) {
	a.Server.AddRoute("GET /api/v1/rooms", api.MakeHandler(a.HandleGetPokerRooms))
	a.Server.AddRoute("GET /api/v1/rooms/{id}", api.MakeHandler(a.HandleGetPokerRoom))
	a.Server.AddRoute("POST /api/v1/rooms", api.MakeHandler(a.HandlePostPokerRooms))
}

// HandleGetPokerRoom @Summary Get all poker rooms
// @Description Get all poker rooms
// @Tags rooms/v1
// @Param id path string true "Room ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} PokerRoomDto
// @Failure 404 {object} errors.ErrorResult
// @Router /api/v1/rooms/{id} [get]
func (a Api) HandleGetPokerRoom(w http.ResponseWriter, r *http.Request) *errors.ErrorResult {
	id := r.PathValue("id")
	parsedId, err := uuid.Parse(id)
	if err != nil || parsedId == uuid.Nil {
		return errors.NewErrorResult(http.StatusBadRequest, "Invalid id", err.Error())
	}
	room, err := a.IPokerRoomRepository.FindByID(parsedId)
	if err != nil {
		return errors.NewErrorResult(
			http.StatusNotFound,
			"Room not found",
			err.Error())
	}
	_ = api.WriteJSON(w, http.StatusOK, ToApi(room))
	return nil
}

// HandleGetPokerRooms @Summary Get all poker rooms
// @Description Get all poker rooms
// @Tags rooms/v1
// @Accept  json
// @Produce  json
// @Success 200 {object} PokerRoomDto[]
// @Failure 500 {object} errors.ErrorResult
// @Router /api/v1/rooms [get]
func (a Api) HandleGetPokerRooms(w http.ResponseWriter, r *http.Request) *errors.ErrorResult {
	allRooms, err := a.IPokerRoomRepository.FindAll()
	if err != nil {
		return errors.NewErrorResult(
			http.StatusInternalServerError,
			"Error getting rooms",
			err.Error())
	}
	_ = api.WriteJSON(w, http.StatusOK, ToApis(allRooms))
	return nil
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
func (a Api) HandlePostPokerRooms(w http.ResponseWriter, r *http.Request) *errors.ErrorResult {
	createRoomReq := new(CreateRoomRequest)
	if err := json.NewDecoder(r.Body).Decode(createRoomReq); err != nil {
		return errors.NewErrorResult(
			http.StatusBadRequest,
			"Invalid request",
			err.Error())
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	room := &PokerRoom{
		Name:    createRoomReq.Name,
		OwnerId: uuid.New(),
	}
	save, err := a.IPokerRoomRepository.Save(room)
	if err != nil {
		return nil
	}
	response := api.CreateResponse{Id: save.ID()}
	_ = api.WriteJSON(w, http.StatusCreated, response)
	return nil
}
