package identity

import (
	contracts "scrum-daddy-be/identitycontracts"
)

type Contract struct {
	IdentityContainer *Container
}

func NewIdentityContracts(c *Container) contracts.IIdentityContracts {
	return &Contract{IdentityContainer: c}
}

func (uc *Contract) CreateGuest(user contracts.CreateQuestUserRequest) (contracts.CreateQuestUserResponse, error) {
	return CreateGuestUser(uc.IdentityContainer, &user)
}

func (uc *Contract) FindById(id string) (contracts.UserResponse, error) {
	return contracts.UserResponse{}, nil
}
