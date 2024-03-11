package pokerplanning

import (
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/identitycontracts"
	"scrum-daddy-be/pokerplanning/abstractions"
	"scrum-daddy-be/pokerplanning/infrastructure"
)

type Container struct {
	Db     *db.Database
	Server *api.Server
	abstractions.IPokerRoomRepository
	IdentityContracts identitycontracts.IIdentityContracts
}

func NewPokerPlanningContainer(
	db *db.Database,
	s *api.Server,
	roomRepo abstractions.IPokerRoomRepository,
	identityContracts identitycontracts.IIdentityContracts) *Container {
	return &Container{
		Db:                   db,
		Server:               s,
		IPokerRoomRepository: roomRepo,
		IdentityContracts:    identityContracts,
	}
}

func NewPokerRoomRepository(db *db.Database) abstractions.IPokerRoomRepository {
	return &infrastructure.PokerRoomRepository{Db: db}
}
