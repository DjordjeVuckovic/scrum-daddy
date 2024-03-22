package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"scrum-daddy-be/common/results"
	"scrum-daddy-be/pokerplanning/utils"
)

type VotingStrategy string

const (
	Fibonacci         VotingStrategy = "fibonacci"
	ModifiedFibonacci VotingStrategy = "modifiedFibonacci"
	TShirt            VotingStrategy = "tshirt"
	Sequential        VotingStrategy = "sequential"
	PowersOfTwo       VotingStrategy = "powersOfTwo"
)

type PokerRoom struct {
	gorm.Model
	ID             uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	SecondaryId    int       `gorm:"autoIncrement;index"`
	Name           string
	OwnerId        uuid.UUID      `gorm:"type:uuid;index"`
	VotingStrategy VotingStrategy `gorm:"type:varchar(25);default:'fibonacci'"`
	IsAllReveal    bool           `gorm:"type:boolean;default:true"`
	AutoReveal     bool           `gorm:"type:boolean;default:false"`
	ShowAverage    bool           `gorm:"type:boolean;default:true"`
}

func (PokerRoom) TableName() string {
	return utils.GetPokerPlanningSchemaName() + ".poker_rooms"
}

func NewPokerRoom(name string, ownerId uuid.UUID) (*PokerRoom, error) {
	room := PokerRoom{
		Name:    name,
		OwnerId: ownerId,
	}
	err := room.validate()
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (pr PokerRoom) validate() error {
	if pr.Name == "" {
		return results.ValidationError(
			"Error validating poker room",
			"Poker room name is required",
			results.ValidationErrType,
		)
	}
	return nil
}
