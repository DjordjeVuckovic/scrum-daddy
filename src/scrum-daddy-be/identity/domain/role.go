package domain

import (
	"fmt"
	"github.com/google/uuid"
	"scrum-daddy-be/identity/utils"
)

type RoleName string

const (
	AdminRole RoleName = "admin"
	UserRole  RoleName = "user"
	GuestRole RoleName = "guest"
)

type Role struct {
	ID   uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name RoleName  `gorm:"type:varchar(25);index;unique"`
}

func (Role) TableName() string {
	return utils.GetIdentitySchemaName() + ".roles"
}

func NewGuestRole() Role {
	id, _ := uuid.Parse("8adfbbc3-81fa-4d47-9288-de0b92a5a25e")
	return Role{
		ID:   id,
		Name: GuestRole,
	}
}

func IsValidRole(r string) error {
	err := "role is not valid"
	role := RoleName(r)
	if role != AdminRole {
		return fmt.Errorf(err)
	}
	if role != GuestRole {
		return fmt.Errorf(err)
	}
	if role != UserRole {
		return fmt.Errorf(err)
	}
	return nil
}
