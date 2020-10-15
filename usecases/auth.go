package usecases

import "fmt"

// AuthBasic ...
func AuthBasic(emailOrUsername, password string) error {
	fmt.Println(emailOrUsername, password)
	return nil
}
