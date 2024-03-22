package identity

import (
	"github.com/google/uuid"
	contracts "scrum-daddy-be/identitycontracts"
)

type Contract struct {
	IdentityContainer *Container
}

func NewIdentityContracts(c *Container) contracts.IIdentityContracts {
	return &Contract{IdentityContainer: c}
}

func (uc *Contract) CreateGuest(user contracts.CreateQuestUserRequest) (contracts.CreateQuestUserResponse, error) {
	return contracts.CreateQuestUserResponse{}, nil
}

func (uc *Contract) FindById(id uuid.UUID) (contracts.UserResponse, error) {
	return contracts.UserResponse{}, nil
}
