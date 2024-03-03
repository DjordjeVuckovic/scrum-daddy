package rooms

import (
	"github.com/google/uuid"
	"log"
	"scrum-daddy-be/common/db"
)

type Reader interface {
	FindByID(id uuid.UUID) (*PokerRoom, error)
	FindAll() ([]*PokerRoom, error)
}

type Writer interface {
	Save(room *PokerRoom) (uuid.UUID, error)
}

type IPokerRoomRepository interface {
	Reader
	Writer
}

type PokerRoomRepository struct {
	db *db.Database
}

func NewPokerRoomRepository(db *db.Database) IPokerRoomRepository {
	return &PokerRoomRepository{db: db}
}

func (r *PokerRoomRepository) FindByID(id uuid.UUID) (*PokerRoom, error) {
	var pokerRoom PokerRoom
	err := r.db.GetDB().First(&pokerRoom, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &pokerRoom, nil
}

func (r *PokerRoomRepository) Save(room *PokerRoom) (uuid.UUID, error) {
	result := r.db.GetDB().Create(room)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return room.ID, nil
}

func (r *PokerRoomRepository) FindAll() ([]*PokerRoom, error) {
	var rooms []*PokerRoom
	result := r.db.GetDB().Find(&rooms)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}
	return rooms, nil
}
