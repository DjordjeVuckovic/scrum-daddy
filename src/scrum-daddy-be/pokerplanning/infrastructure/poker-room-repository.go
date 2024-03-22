package infrastructure

import (
	"context"
	"github.com/google/uuid"
	"log"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/pokerplanning/domain"
)

type PokerRoomRepository struct {
	Uow db.IUnitOfWork
}

func (r *PokerRoomRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.PokerRoom, error) {
	var pokerRoom domain.PokerRoom
	err := r.Uow.GetConnection().WithContext(ctx).First(&pokerRoom, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &pokerRoom, nil
}

func (r *PokerRoomRepository) Save(ctx context.Context, room *domain.PokerRoom) (uuid.UUID, error) {
	result := r.Uow.GetConnection().WithContext(ctx).Create(room)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return room.ID, nil
}

func (r *PokerRoomRepository) FindAll(ctx context.Context) ([]*domain.PokerRoom, error) {
	var rooms []*domain.PokerRoom
	result := r.Uow.GetConnection().WithContext(ctx).Find(&rooms)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}
	return rooms, nil
}

func (r *PokerRoomRepository) FindBySecondaryID(ctx context.Context, id int) (*domain.PokerRoom, error) {
	var pokerRoom domain.PokerRoom
	err := r.Uow.GetConnection().WithContext(ctx).First(&pokerRoom, "secondary_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &pokerRoom, nil
}
