package infrastructure

import (
	"github.com/google/uuid"
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
		panic(err)
	}
	return &UserRepository{uow: uow}
}

func NewUserRepository(dba *db.Database) abstractions.IUserRepository {
	return UserRepositoryFactory(dba, db.UowFactory)
}

func (r *UserRepository) CreateUser(user *domain.User) (uuid.UUID, error) {
	result := r.uow.GetConnection().Db.Create(user)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return user.ID, nil
}

func (r *UserRepository) FindById(id uuid.UUID) (*domain.User, error) {
	var usr domain.User
	err := r.uow.GetConnection().Db.First(&usr, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &usr, nil
}
