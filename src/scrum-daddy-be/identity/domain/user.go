package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"scrum-daddy-be/identity/utils"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Username string
	Password string
	Email    string
	Roles    []Role `gorm:"type:json"`
}

func (User) TableName() string {
	return utils.GetIdentitySchemaName() + ".users"
}
