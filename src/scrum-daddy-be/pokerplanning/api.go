package pokerplanning

import (
	"github.com/google/uuid"
	"net/http"
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/errors"
	"scrum-daddy-be/pokerplanning/dto"
)

func Main(c *Container) {
	MigrateDb(c.Db)
	AddPokerRooms(c)
}

func AddPokerRooms(c *Container) {
	c.Server.AddRoute("GET /api/v1/rooms", api.MakeHandler(c.HandleGetPokerRooms))

	c.Server.AddRoute("POST /api/v1/rooms", api.MakeHandler(c.HandlePostPokerRooms))

	c.Server.AddRoute("GET /api/v1/rooms/{id}", api.MakeHandler(c.HandleGetPokerRoom))
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
func (c Container) HandleGetPokerRoom(w http.ResponseWriter, r *http.Request) *errors.ErrorResult {
	id := r.PathValue("id")
	parsedId, err := uuid.Parse(id)
	if err != nil || parsedId == uuid.Nil {
		return errors.NewErrorResult(http.StatusBadRequest, "Invalid id", err.Error())
	}
	room, err := c.IPokerRoomRepository.FindByID(parsedId)
	if err != nil {
		return errors.NewErrorResult(
			http.StatusNotFound,
			"Room not found",
			err.Error())
	}
	_ = api.WriteJSON(w, http.StatusOK, dto.ToApi(room))
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
func (c Container) HandleGetPokerRooms(w http.ResponseWriter, r *http.Request) *errors.ErrorResult {
	allRooms, err := c.IPokerRoomRepository.FindAll()
	if err != nil {
		return errors.NewErrorResult(
			http.StatusInternalServerError,
			"Error getting rooms",
			err.Error())
	}
	_ = api.WriteJSON(w, http.StatusOK, dto.ToApis(allRooms))
	return nil
}
