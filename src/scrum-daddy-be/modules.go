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

func CreateModules(s *api.Server, db *db.Database) Modules {
	// identity
	identityContainer := identity.NewIdentityContainer(
		db,
		s,
	)

	// identity contracts
	identityContracts := identity.NewIdentityContracts(identityContainer)

	// pokerPlanning
	pokerPlanningContainer := pokerplanning.NewPokerPlanningContainer(
		db,
		s,
		identityContracts,
	)

	return Modules{
		PokerPlanning: pokerPlanningContainer,
		Identity:      identityContainer,
	}
}
