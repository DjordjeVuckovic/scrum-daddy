package identity

import (
	"github.com/google/uuid"
	"scrum-daddy-be/identity/common"
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
	return common.GetIdentitySchemaName() + ".roles"
}
