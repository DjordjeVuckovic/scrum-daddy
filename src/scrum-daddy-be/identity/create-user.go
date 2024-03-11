package identity

import (
	"log"
	contracts "scrum-daddy-be/identitycontracts"
)

func CreateGuestUser(c *Container, userReq *contracts.CreateQuestUserRequest) (contracts.CreateQuestUserResponse, error) {
	role, err := c.IRoleRepository.FindByName(GuestRole)
	if err != nil {
		log.Fatal(err)
		return contracts.CreateQuestUserResponse{}, err
	}
	user := &User{
		Username: userReq.Username,
		Roles: []Role{
			role,
		},
	}
	userId, err := c.IUserRepository.CreateUser(user)
	if err != nil {
		log.Fatal(err)
		return contracts.CreateQuestUserResponse{}, err
	}
	userResp := contracts.CreateQuestUserResponse{
		ID:       userId,
		Username: user.Username,
	}
	return userResp, nil
}
