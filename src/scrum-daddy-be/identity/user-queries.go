package identity

import (
	"github.com/google/uuid"
	"scrum-daddy-be/identity/abstractions"
	"scrum-daddy-be/identity/dto"
)

func FindById(
	repository abstractions.IUserRepository,
	id uuid.UUID) (dto.UserResponse, error) {
	user, err := repository.FindById(id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.ToApi(*user), nil
}
