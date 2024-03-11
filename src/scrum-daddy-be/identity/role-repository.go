package identity

import (
	"scrum-daddy-be/common/db"
)

type Reader interface {
	FindByName(name RoleName) (Role, error)
}
type IRoleRepository interface {
	Reader
}

type RoleRepository struct {
	DB *db.Database
}

func NewRoleRepository(Db *db.Database) IRoleRepository {
	return &RoleRepository{DB: Db}
}

func (r *RoleRepository) FindByName(name RoleName) (Role, error) {
	var role Role
	result := r.DB.GetDB().Where("name = ?", name).First(&role)
	if result.Error != nil {
		return Role{}, result.Error
	}
	return role, nil
}
