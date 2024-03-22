package abstractions

import (
	"context"
	"github.com/google/uuid"
	"scrum-daddy-be/pokerplanning/domain"
)

type Reader interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.PokerRoom, error)
	FindBySecondaryID(ctx context.Context, id int) (*domain.PokerRoom, error)
	FindAll(ctx context.Context) ([]*domain.PokerRoom, error)
}

type Writer interface {
	Save(ctx context.Context, room *domain.PokerRoom) (uuid.UUID, error)
}

type IPokerRoomRepository interface {
	Reader
	Writer
}
