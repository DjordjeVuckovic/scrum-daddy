package identity

import (
	"context"
	"log/slog"
	"scrum-daddy-be/identity/abstractions"
	"scrum-daddy-be/identity/domain"
	contracts "scrum-daddy-be/identitycontracts"
)

func CreateGuestUser(
	ctx context.Context,
	usersRepo abstractions.IUserRepository,
	userReq *contracts.CreateQuestUserRequest) (contracts.CreateQuestUserResponse, error) {
	user := &domain.User{
		Username: userReq.Username,
		Roles: []domain.Role{
			domain.NewGuestRole(),
		},
	}
	userId, err := usersRepo.CreateUser(ctx, user)
	if err != nil {
		slog.Error("Error while creating user", "err", err)
		return contracts.CreateQuestUserResponse{}, err
	}
	userResp := contracts.CreateQuestUserResponse{
		ID:       userId,
		Username: user.Username,
	}
	return userResp, nil
}
