package identity

import (
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/db"
)

type Container struct {
	Db     *db.Database
	Server *api.Server
}

func Main(c *Container) {
	MigrateDb(c.Db)
}
