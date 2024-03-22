package identitycontracts

import "github.com/google/uuid"

type IIdentityContracts interface {
	IUserContracts
}

type IUserContracts interface {
	CreateGuest(user CreateQuestUserRequest) (CreateQuestUserResponse, error)
	FindById(id uuid.UUID) (UserResponse, error)
}
