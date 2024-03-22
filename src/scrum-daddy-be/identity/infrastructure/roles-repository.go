package infrastructure

import (
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/identity/abstractions"
	"scrum-daddy-be/identity/domain"
)

type RoleRepository struct {
	Uow db.IUnitOfWork
}

type RoleRepositoryFactoryFn func(db *db.Database, uowFn db.UowFactoryFn) abstractions.IRoleRepository

func RoleRepositoryFactory(database *db.Database, uowFn db.UowFactoryFn) abstractions.IRoleRepository {
	uow, err := uowFn(database)
	if err != nil {
		panic(err)
	}
	return &RoleRepository{Uow: uow}
}

func NewRoleRepository(database *db.Database) abstractions.IRoleRepository {
	return RoleRepositoryFactory(database, db.UowFactory)
}

func (r *RoleRepository) FindByName(name domain.RoleName) (domain.Role, error) {
	var role domain.Role
	result := r.Uow.GetConnection().Db.Where("name = ?", name).First(&role)
	if result.Error != nil {
		return domain.Role{}, result.Error
	}
	return role, nil
}
