package dto

import (
	"github.com/google/uuid"
	"scrum-daddy-be/pokerplanning/types"
)

type PokerRoomDto struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	OwnerId uuid.UUID `json:"ownerId"`
}

func ToApi(pokerRoom *types.PokerRoom) PokerRoomDto {
	return PokerRoomDto{
		ID:      pokerRoom.ID,
		Name:    pokerRoom.Name,
		OwnerId: pokerRoom.OwnerId,
	}
}

func ToApis(pokerRooms []*types.PokerRoom) []PokerRoomDto {
	pokerRoomDtos := make([]PokerRoomDto, len(pokerRooms))
	for i, pokerRoom := range pokerRooms {
		pokerRoomDtos[i] = ToApi(pokerRoom)
	}
	return pokerRoomDtos
}
