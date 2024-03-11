package identity

import (
	"github.com/google/uuid"
	"scrum-daddy-be/common/db"
)

type IUserRepository interface {
	CreateUser(user *User) (uuid.UUID, error)
}

type UserRepository struct {
	DB *db.Database
}

func NewUserRepository(db *db.Database) IUserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *User) (uuid.UUID, error) {
	result := r.DB.GetDB().Create(user)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return user.ID, nil
}
