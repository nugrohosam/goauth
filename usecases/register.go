package usecases

import (
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	userRepo "github.com/nugrohosam/gosampleapi/repositories/user"
)

// RegisterBasic ...
func RegisterBasic(name, username, email, password string) error {
	hashedPassword := helpers.MakeHash(password)
	if _, err := userRepo.Create(name, username, email, hashedPassword); err != nil {
		return err
	}

	return nil
}
