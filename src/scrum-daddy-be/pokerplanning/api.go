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

	c.Server.AddRoute("POST /api/v1/rooms", api.MakeHandler(c.HandlePostPokerRooms))

	c.Server.AddRoute("GET /api/v1/rooms/{id}", api.MakeHandler(c.HandleGetPokerRoom))

	c.Server.AddRoute(
		"GET /api/v1/rooms/secondary/{id}",
		api.MakeHandler(c.HandleGetSequentialPokerRoom))

	//hub := NewRoomManager()
	//go hub.listen()
	//http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	//	hub.serveWs(w, r)
	//})
}
