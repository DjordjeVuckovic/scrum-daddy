package pokerplanning

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/common/results"
	"scrum-daddy-be/contracts/identitycontracts"
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

// HandlePostPokerRoom @Summary Create new poker roomId
// @Detail Create new poker roomId
// @Tags rooms
// @Accept  json
// @Produce  json
// @Param   requestBody  body      CreateRoomRequest  true  "Create Room Request"
// @Success 200 {object} api.CreateResponse
// @Failure 400 {object} results.ErrorResult
// @Router /api/v1/rooms [post]
func (c Container) HandlePostPokerRoom(w http.ResponseWriter, r *http.Request) *results.ErrorResult {
	createRoomReq := new(CreateRoomRequest)
	if err := json.NewDecoder(r.Body).Decode(createRoomReq); err != nil {
		return results.NewErrorResult(
			http.StatusBadRequest,
			"Invalid request",
			err.Error())
	}
	log.Println("Create roomId request", *createRoomReq)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	repo := NewPokerRoomRepository(c.Db)
	roomId, err := c.CreateRoom(r.Context(), repo, createRoomReq)

	if err != nil {
		var errResult *results.ErrorResult
		errors.As(err, &errResult)
		return errResult
	}
	response := api.CreateResponse{Id: roomId}
	_ = api.WriteJSON(w, http.StatusCreated, response)
	return nil
}

func (c Container) CreateRoom(
	ctx context.Context,
	repo abstractions.IPokerRoomRepository,
	request *CreateRoomRequest) (uuid.UUID, error) {
	room, err := domain.NewPokerRoom(request.Name, request.OwnerId)
	if err != nil {
		return uuid.Nil, err
	}
	id, err := repo.Save(ctx, room)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

type CreateRoomWithUserRequest struct {
	Name         string         `json:"name"`
	VotingSystem string         `json:"votingSystem"`
	IsAllReveal  bool           `json:"isAllReveal"`
	AutoReveal   bool           `json:"autoReveal"`
	ShowAverage  bool           `json:"showAverage"`
	User         UserCreateRoom `json:"user"`
}

type UserCreateRoom struct {
	Username string `json:"username"`
}

// HandlePostPokerRoomWithUser @Summary Create new poker roomId
// @Detail Create new poker roomId
// @Tags rooms
// @Accept  json
// @Produce  json
// @Param   requestBody  body      CreateRoomWithUserRequest  true  "Create Room Request"
// @Success 200 {object} api.CreateResponse
// @Failure 400 {object} results.ErrorResult
// @Router /api/v1/rooms/with-user [post]
func (c Container) HandlePostPokerRoomWithUser(w http.ResponseWriter, r *http.Request) *results.ErrorResult {
	createRoomReq := new(CreateRoomWithUserRequest)
	if err := json.NewDecoder(r.Body).Decode(createRoomReq); err != nil {
		return results.NewErrorResult(
			http.StatusBadRequest,
			"Invalid request",
			err.Error())
	}
	log.Println("Create roomId request", *createRoomReq)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	repo, uow := NewTransactionalPokerRoomRepository(c.Db)
	uowCtx := context.WithValue(r.Context(), "uow", uow)
	roomId, err := c.CreateRoomWithUser(uowCtx, uow, repo, createRoomReq)

	if err != nil {
		var errResult *results.ErrorResult
		errors.As(err, &errResult)
		return errResult
	}
	response := api.CreateResponse{Id: roomId}
	_ = api.WriteJSON(w, http.StatusCreated, response)
	return nil
}

func (c Container) CreateRoomWithUser(
	ctx context.Context,
	uow db.IUnitOfWork,
	repo abstractions.IPokerRoomRepository,
	request *CreateRoomWithUserRequest) (uuid.UUID, error) {
	user := identitycontracts.CreateQuestUserRequest{
		Username: request.User.Username,
	}
	createdUser, err := c.IdentityContracts.CreateGuest(ctx, user)
	if err != nil {
		_ = uow.Rollback()
		return uuid.Nil, err
	}

	room, err := domain.NewPokerRoom(request.Name, createdUser.ID)
	if err != nil {
		return uuid.Nil, err
	}
	id, err := repo.Save(ctx, room)
	if err != nil {
		_ = uow.Rollback()
		return uuid.Nil, err
	}

	_ = uow.Commit()

	return id, nil
}
