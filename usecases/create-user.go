package usecases

import (
	"fmt"

	repo "github.com/nugrohosam/gosampleapi/repositories/user"
)

// CreateUser is use like application
func CreateUser(name string) {
	user := repo.Create(name)
	fmt.Println(user)
}
