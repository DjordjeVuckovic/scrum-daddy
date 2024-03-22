package pokerplanning

import (
	"log/slog"
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/identitycontracts"
	"scrum-daddy-be/pokerplanning/abstractions"
	"scrum-daddy-be/pokerplanning/infrastructure"
)

type Container struct {
	Db                *db.Database
	Server            *api.Server
	IdentityContracts identitycontracts.IIdentityContracts
}

func NewPokerPlanningContainer(
	db *db.Database,
	s *api.Server,
	identityContracts identitycontracts.IIdentityContracts) *Container {
	return &Container{
		Db:                db,
		Server:            s,
		IdentityContracts: identityContracts,
	}
}

type PokerRoomRepositoryFactoryFn func(db *db.Database, uowFn db.UowFactoryFn) abstractions.IPokerRoomRepository

func PokerRoomRepositoryFactory(database *db.Database, uowFn db.UowFactoryFn) abstractions.IPokerRoomRepository {
	uow, err := uowFn(database)
	if err != nil {
		slog.Error("Error creating unit of work", "error", err)
	}
	return &infrastructure.PokerRoomRepository{Uow: uow}
}

func PokerRoomRepositoryTransactionalFactory(database *db.Database, uowFn db.UowFactoryFn) (abstractions.IPokerRoomRepository, db.IUnitOfWork) {
	uow, err := uowFn(database)
	if err != nil {
		slog.Error("Error creating unit of work", "error", err)
	}
	return &infrastructure.PokerRoomRepository{Uow: uow}, uow
}

func NewPokerRoomRepository(database *db.Database) abstractions.IPokerRoomRepository {
	return PokerRoomRepositoryFactory(database, db.UowFactory)
}

func NewTransactionalPokerRoomRepository(database *db.Database) (abstractions.IPokerRoomRepository, db.IUnitOfWork) {
	return PokerRoomRepositoryTransactionalFactory(database, db.UowTransactionalFactory)
}
