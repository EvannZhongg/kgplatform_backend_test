package users

import (
	"kgplatform-backend/api/users"
	userLogic "kgplatform-backend/internal/logic/users"
)

type ControllerV1 struct {
	users *userLogic.Users
}

func NewV1() users.IUsersV1 {
	return &ControllerV1{
		users: userLogic.New(),
	}
}
