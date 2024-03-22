package abstractions

import "scrum-daddy-be/identity/domain"

type IRoleRepositoryReader interface {
	FindByName(name domain.RoleName) (domain.Role, error)
}
type IRoleRepository interface {
	IRoleRepositoryReader
}
