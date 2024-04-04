package infrastructure

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/identity/abstractions"
	"scrum-daddy-be/identity/domain"
)

type UserRepository struct {
	uow db.IUnitOfWork
}

func UserRepositoryFactory(database *db.Database, uowFn db.UowFactoryFn) abstractions.IUserRepository {
	uow, err := uowFn(database)
	if err != nil {
		slog.Error("Err")
	}
	return &UserRepository{uow: uow}
}

func NewUserRepository(dba *db.Database) abstractions.IUserRepository {
	return UserRepositoryFactory(dba, db.UowFactory)
}
func NewUserRepositoryWithUow(uow db.IUnitOfWork) abstractions.IUserRepository {
	return &UserRepository{uow: uow}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) (uuid.UUID, error) {
	result := r.uow.GetConnection().WithContext(ctx).Create(user)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return user.ID, nil
}

func (r *UserRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var usr domain.User
	err := r.uow.GetConnection().WithContext(ctx).First(&usr, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &usr, nil
}
