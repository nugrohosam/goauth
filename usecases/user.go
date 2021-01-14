package usecases

import (
	userRepo "github.com/nugrohosam/gosampleapi/repositories/user"
)

// ShowUser ..
func ShowUser(ID string) userRepo.User {
	user := userRepo.Find(ID)
	return user
}