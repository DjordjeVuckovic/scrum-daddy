package infrastructure

import (
	"github.com/google/uuid"
	"log"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/pokerplanning/types"
)

type PokerRoomRepository struct {
	Db *db.Database
}

func (r *PokerRoomRepository) FindByID(id uuid.UUID) (*types.PokerRoom, error) {
	var pokerRoom types.PokerRoom
	err := r.Db.GetDB().First(&pokerRoom, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &pokerRoom, nil
}

func (r *PokerRoomRepository) Save(room *types.PokerRoom) (uuid.UUID, error) {
	result := r.Db.GetDB().Create(room)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return room.ID, nil
}

func (r *PokerRoomRepository) FindAll() ([]*types.PokerRoom, error) {
	var rooms []*types.PokerRoom
	result := r.Db.GetDB().Find(&rooms)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}
	return rooms, nil
}
