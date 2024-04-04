package identitycontracts

import (
	"context"
	"github.com/google/uuid"
)

type IIdentityContracts interface {
	IUserContracts
}

type IUserContracts interface {
	CreateGuest(ctx context.Context, user CreateQuestUserRequest) (CreateQuestUserResponse, error)
	FindById(id uuid.UUID) (UserResponse, error)
}
