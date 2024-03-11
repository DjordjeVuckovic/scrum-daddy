package identity

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
