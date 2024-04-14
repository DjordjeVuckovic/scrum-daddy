package pokerplanning

import (
	"scrum-daddy-be/common/api"
)

func Main(c *Container) {
	MigrateDb(c.Db)
	AddPokerRooms(c)
}

func AddPokerRooms(c *Container) {
	c.Server.AddRoute("GET /api/v1/rooms", api.MakeHandler(c.HandleGetPokerRooms))

	c.Server.AddRoute("POST /api/v1/rooms", api.MakeHandler(c.HandlePostPokerRoom))

	c.Server.AddRoute("POST /api/v1/rooms/with-user", api.MakeHandler(c.HandlePostPokerRoomWithUser))

	c.Server.AddRoute("GET /api/v1/rooms/{id}", api.MakeHandler(c.HandleGetPokerRoom))

	c.Server.AddRoute(
		"GET /api/v1/rooms/secondary/{id}",
		api.MakeHandler(c.HandleGetPokerRoomBySecondaryId))

	go c.Hub.listen()
	c.Server.GetMux().HandleFunc("/ws/{roomId}", c.Hub.serveWs)
}
