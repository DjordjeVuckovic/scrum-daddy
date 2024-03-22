package dto

import "scrum-daddy-be/identity/domain"

type RoleResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserResponse struct {
	ID       string         `json:"id"`
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Roles    []RoleResponse `json:"roles"`
}

func ToApi(user domain.User) UserResponse {
	return UserResponse{
		ID:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
		Roles:    []RoleResponse{},
	}
}
