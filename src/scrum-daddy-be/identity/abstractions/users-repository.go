package abstractions

import (
	"github.com/google/uuid"
	"scrum-daddy-be/identity/domain"
)

type IUserRepositoryWriter interface {
	CreateUser(user *domain.User) (uuid.UUID, error)
}

type IUserRepositoryReader interface {
	FindById(id uuid.UUID) (*domain.User, error)
}

type IUserRepository interface {
	IUserRepositoryWriter
	IUserRepositoryReader
}
