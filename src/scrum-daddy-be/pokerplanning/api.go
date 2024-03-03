package pokerplanning

import (
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/pokerplanning/rooms"
)

func Main(s *api.Server) {
	dbConnection := db.Connect()
	MigrateDb(dbConnection)

	roomRepo := rooms.NewPokerRoomRepository(dbConnection)
	roomsApi := rooms.Api{Db: dbConnection, Server: s, IPokerRoomRepository: roomRepo}
	rooms.AddPokerRooms(roomsApi)

}
