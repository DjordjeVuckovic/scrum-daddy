package identity

import (
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/identity/common"
)

type Container struct {
	Db     *db.Database
	Server *api.Server
	IUserRepository
	IRoleRepository
}

func Main(c *Container) {
	common.MigrateDb(c.Db)
}
