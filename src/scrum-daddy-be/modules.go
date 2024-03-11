package main

import (
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/identity"
	"scrum-daddy-be/pokerplanning"
)

type Modules struct {
	PokerPlanning *pokerplanning.Container
	Identity      *identity.Container
}

func CreateModuleContainers(s *api.Server, db *db.Database) Modules {
	// identity
	usersRepo := identity.NewUserRepository(db)
	rolesRepo := identity.NewRoleRepository(db)
	identityContainer := NewIdentityContainer(db, s, usersRepo, rolesRepo)

	// identity contracts
	identityContracts := identity.NewIdentityContracts(identityContainer)

	// pokerPlanning
	roomRepo := pokerplanning.NewPokerRoomRepository(db)
	pokerPlanningContainer := pokerplanning.NewPokerPlanningContainer(db, s, roomRepo, identityContracts)

	return Modules{
		PokerPlanning: pokerPlanningContainer,
		Identity:      identityContainer,
	}
}

func NewIdentityContainer(
	db *db.Database,
	s *api.Server,
	userRepo identity.IUserRepository,
	roleRepo identity.IRoleRepository) *identity.Container {
	return &identity.Container{
		Db:              db,
		Server:          s,
		IUserRepository: userRepo,
		IRoleRepository: roleRepo,
	}
}
