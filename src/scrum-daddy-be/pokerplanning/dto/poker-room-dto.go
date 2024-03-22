package dto

import (
	"github.com/google/uuid"
	"scrum-daddy-be/pokerplanning/domain"
)

type PokerRoomDto struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	OwnerId uuid.UUID `json:"ownerId"`
}

func ToApi(pokerRoom *domain.PokerRoom) PokerRoomDto {
	return PokerRoomDto{
		ID:      pokerRoom.ID,
		Name:    pokerRoom.Name,
		OwnerId: pokerRoom.OwnerId,
	}
}

func ToApis(pokerRooms []*domain.PokerRoom) []PokerRoomDto {
	pokerRoomDtos := make([]PokerRoomDto, len(pokerRooms))
	for i, pokerRoom := range pokerRooms {
		pokerRoomDtos[i] = ToApi(pokerRoom)
	}
	return pokerRoomDtos
}
