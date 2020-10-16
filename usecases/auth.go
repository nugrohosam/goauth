package usecases

import (
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	userRepo "github.com/nugrohosam/gosampleapi/repositories/user"
)

// AuthBasic ...
func AuthBasic(emailOrUsername, password string) error {
	hashedPassword := helpers.MakeHash(password)
	userRepo.FindByEmailOrUsernameAndPassword(emailOrUsername, hashedPassword)

	return nil
}
