package identity

import (
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/db"
)

func NewIdentityContainer(
	db *db.Database,
	s *api.Server) *Container {
	return &Container{
		Db:     db,
		Server: s,
	}
}
