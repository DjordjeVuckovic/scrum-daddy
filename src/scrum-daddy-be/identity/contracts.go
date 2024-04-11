package identity

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/contracts/identitycontracts"
	"scrum-daddy-be/identity/infrastructure"
)

type Contract struct {
	IdentityContainer *Container
}

func NewIdentityContracts(c *Container) identitycontracts.IIdentityContracts {
	return &Contract{IdentityContainer: c}
}

func (uc *Contract) CreateGuest(ctx context.Context, user identitycontracts.CreateQuestUserRequest) (identitycontracts.CreateQuestUserResponse, error) {
	ctxUow := ctx.Value("uow")
	uow, ok := ctxUow.(db.IUnitOfWork)
	if !ok {
		return identitycontracts.CreateQuestUserResponse{}, fmt.Errorf("error while reading uow key")
	}
	usersRepo := infrastructure.NewUserRepositoryWithUow(uow)
	u, err := CreateGuestUser(ctx, usersRepo, &user)
	if err != nil {
		return identitycontracts.CreateQuestUserResponse{}, err
	}
	return u, nil
}

func (uc *Contract) FindById(id uuid.UUID) (identitycontracts.UserResponse, error) {
	return identitycontracts.UserResponse{}, nil
}
