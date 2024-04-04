package identity

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/identity/infrastructure"
	contracts "scrum-daddy-be/identitycontracts"
)

type Contract struct {
	IdentityContainer *Container
}

func NewIdentityContracts(c *Container) contracts.IIdentityContracts {
	return &Contract{IdentityContainer: c}
}

func (uc *Contract) CreateGuest(ctx context.Context, user contracts.CreateQuestUserRequest) (contracts.CreateQuestUserResponse, error) {
	ctxUow := ctx.Value("uow")
	uow, ok := ctxUow.(db.IUnitOfWork)
	if !ok {
		return contracts.CreateQuestUserResponse{}, fmt.Errorf("error while reading uow key")
	}
	usersRepo := infrastructure.NewUserRepositoryWithUow(uow)
	u, err := CreateGuestUser(ctx, usersRepo, &user)
	if err != nil {
		return contracts.CreateQuestUserResponse{}, err
	}
	return u, nil
}

func (uc *Contract) FindById(id uuid.UUID) (contracts.UserResponse, error) {
	return contracts.UserResponse{}, nil
}
