package identitycontracts

import "github.com/google/uuid"

type CreateQuestUserRequest struct {
	Username string `json:"username"`
}

type CreateQuestUserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type UserResponse struct {
	ID       string         `json:"id"`
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Roles    []RoleResponse `json:"roles"`
}

type RoleResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
