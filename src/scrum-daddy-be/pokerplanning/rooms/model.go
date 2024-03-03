package rooms

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"scrum-daddy-be/pokerplanning/common"
)

type PokerRoom struct {
	gorm.Model
	ID      uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name    string
	OwnerId uuid.UUID `gorm:"type:uuid;index"`
}

func (PokerRoom) TableName() string {
	return common.GetPokerPlanningSchemaName() + ".poker_rooms"
}

func NewPokerRoom(name string, ownerId uuid.UUID) *PokerRoom {
	return &PokerRoom{
		Name:    name,
		OwnerId: ownerId,
	}
}
