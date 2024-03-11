package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"scrum-daddy-be/pokerplanning/utils"
)

type VotingSystem string

const (
	Fibonacci         VotingSystem = "fibonacci"
	ModifiedFibonacci VotingSystem = "modifiedFibonacci"
	TShirt            VotingSystem = "tshirt"
	Sequential        VotingSystem = "sequential"
	PowersOfTwo       VotingSystem = "powersOfTwo"
)

type PokerRoom struct {
	gorm.Model
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name         string
	OwnerId      uuid.UUID    `gorm:"type:uuid;index"`
	VotingSystem VotingSystem `gorm:"type:varchar(25);default:'fibonacci'"`
	IsAllReveal  bool         `gorm:"type:boolean;default:true"`
	AutoReveal   bool         `gorm:"type:boolean;default:false"`
	ShowAverage  bool         `gorm:"type:boolean;default:true"`
}

func (PokerRoom) TableName() string {
	return utils.GetPokerPlanningSchemaName() + ".poker_rooms"
}

func NewPokerRoom(name string, ownerId uuid.UUID) *PokerRoom {
	return &PokerRoom{
		Name:    name,
		OwnerId: ownerId,
	}
}
