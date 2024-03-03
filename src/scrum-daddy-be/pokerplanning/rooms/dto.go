package rooms

import (
	"github.com/google/uuid"
)

type PokerRoomDto struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	OwnerId uuid.UUID `json:"ownerId"`
}

func ToApi(pokerRoom *PokerRoom) PokerRoomDto {
	return PokerRoomDto{
		ID:      pokerRoom.ID,
		Name:    pokerRoom.Name,
		OwnerId: pokerRoom.OwnerId,
	}
}

func ToApis(pokerRooms []*PokerRoom) []PokerRoomDto {
	pokerRoomDtos := make([]PokerRoomDto, len(pokerRooms))
	for i, pokerRoom := range pokerRooms {
		pokerRoomDtos[i] = ToApi(pokerRoom)
	}
	return pokerRoomDtos
}

type CreateRoomRequest struct {
	Name string `json:"name"`
}
