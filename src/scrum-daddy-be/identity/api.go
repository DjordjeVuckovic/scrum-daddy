package identity

import (
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/identity/abstractions"
)

type Container struct {
	Db              *db.Database
	Server          *api.Server
	IRoleRepository abstractions.IRoleRepository
}

func Main(c *Container) {
	MigrateDb(c.Db)
}
