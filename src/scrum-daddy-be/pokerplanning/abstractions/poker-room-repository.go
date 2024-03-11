package abstractions

import (
	"github.com/google/uuid"
	"scrum-daddy-be/pokerplanning/types"
)

type Reader interface {
	FindByID(id uuid.UUID) (*types.PokerRoom, error)
	FindAll() ([]*types.PokerRoom, error)
}

type Writer interface {
	Save(room *types.PokerRoom) (uuid.UUID, error)
}

type IPokerRoomRepository interface {
	Reader
	Writer
}
