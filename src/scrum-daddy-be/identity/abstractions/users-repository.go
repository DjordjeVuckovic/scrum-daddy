package abstractions

import (
	"context"
	"github.com/google/uuid"
	"scrum-daddy-be/identity/domain"
)

type IUserRepositoryWriter interface {
	CreateUser(ctx context.Context, user *domain.User) (uuid.UUID, error)
}

type IUserRepositoryReader interface {
	FindById(ctx context.Context, id uuid.UUID) (*domain.User, error)
}

type IUserRepository interface {
	IUserRepositoryWriter
	IUserRepositoryReader
}
