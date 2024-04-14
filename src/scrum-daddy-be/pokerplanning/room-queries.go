package pokerplanning

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/results"
	"scrum-daddy-be/common/utils"
	"scrum-daddy-be/pokerplanning/abstractions"
	"scrum-daddy-be/pokerplanning/domain"
	"scrum-daddy-be/pokerplanning/dto"
)

// HandleGetPokerRoom @Summary Get all poker rooms
// @Detail Get all poker rooms
// @Tags rooms
// @Param id path string true "Room ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.PokerRoomDto
// @Failure 404 {object} results.ErrorResult
// @Router /api/v1/rooms/{id} [get]
func (c Container) HandleGetPokerRoom(w http.ResponseWriter, r *http.Request) *results.ErrorResult {
	id := r.PathValue("id")
	parsedId, err := uuid.Parse(id)
	if err != nil || parsedId == uuid.Nil {
		return results.NewErrorResult(http.StatusBadRequest, "Invalid id", err.Error())
	}

	repo := NewPokerRoomRepository(c.Db)

	room, err := repo.FindByID(r.Context(), parsedId)

	if err != nil {
		return results.NewErrorResult(
			http.StatusNotFound,
			"Room not found",
			err.Error())
	}
	_ = api.WriteJSON(w, http.StatusOK, dto.ToApi(room))
	return nil
}

// HandleGetPokerRoomBySecondaryId @Tags rooms
// @Param id path int true "Room Secondary ID"
// @Tags rooms
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.PokerRoomDto
// @Failure 404 {object} results.ErrorResult
// @Router /api/v1/rooms/secondary/{id} [get]
func (c Container) HandleGetPokerRoomBySecondaryId(w http.ResponseWriter, r *http.Request) *results.ErrorResult {
	id := r.PathValue("id")
	parsedId, err := utils.ParseToInt(id)
	if err != nil {
		return results.NewErrorResult(http.StatusBadRequest, "Invalid id", "Cannot parse route parameter")
	}

	repo := NewPokerRoomRepository(c.Db)
	room, errResult := GetPokerRoomBySecondaryId(
		r.Context(),
		repo,
		parsedId,
	)
	if errResult != nil {
		return errResult
	}
	_ = api.WriteJSON(w, http.StatusOK, dto.ToApi(room))
	return nil
}

func GetPokerRoomBySecondaryId(ctx context.Context, repo abstractions.IPokerRoomRepository, id int) (*domain.PokerRoom, *results.ErrorResult) {

	room, err := repo.FindBySecondaryID(ctx, id)

	if err != nil {
		return nil, results.NewErrorResult(
			http.StatusNotFound,
			"Room not found",
			err.Error())
	}
	return room, nil
}

// HandleGetPokerRooms @Summary Get all poker rooms
// @Detail Get all poker rooms
// @Tags rooms
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} dto.PokerRoomDto[]
// @Router /api/v1/rooms [get]
func (c Container) HandleGetPokerRooms(w http.ResponseWriter, r *http.Request) *results.ErrorResult {
	repo := NewPokerRoomRepository(c.Db)
	allRooms, err := repo.FindAll(r.Context())
	if err != nil {
		return results.NewErrorResult(
			http.StatusInternalServerError,
			"Error getting rooms",
			err.Error())
	}
	_ = api.WriteJSON(w, http.StatusOK, dto.ToApis(allRooms))
	return nil
}
