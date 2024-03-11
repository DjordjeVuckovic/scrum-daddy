package identitycontracts

type IIdentityContracts interface {
	IUserContracts
}

type IUserContracts interface {
	CreateGuest(user CreateQuestUserRequest) (CreateQuestUserResponse, error)
	FindById(id string) (UserResponse, error)
}
