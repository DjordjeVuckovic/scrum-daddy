package identity

import (
	"log"
	"scrum-daddy-be/identity/abstractions"
	"scrum-daddy-be/identity/domain"
	contracts "scrum-daddy-be/identitycontracts"
)

func CreateGuestUser(
	usersRepo abstractions.IUserRepository,
	rolesRepo abstractions.IRoleRepository,
	userReq *contracts.CreateQuestUserRequest) (contracts.CreateQuestUserResponse, error) {
	role, err := rolesRepo.FindByName(domain.GuestRole)
	if err != nil {
		log.Fatal(err)
		return contracts.CreateQuestUserResponse{}, err
	}
	user := &domain.User{
		Username: userReq.Username,
		Roles: []domain.Role{
			role,
		},
	}
	userId, err := usersRepo.CreateUser(user)
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
