package v1

// AuthLogin ...
type AuthLogin struct {
	EmailOrUsername string `json:"emailOrUsername" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
}

// AuthRegister ...
type AuthRegister struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
