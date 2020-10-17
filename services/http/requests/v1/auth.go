package v1

// AuthLogin ...
type AuthLogin struct {
	EmailOrUsername string `json:"emailOrUsername" validate:"required,ascii"`
	Password        string `json:"password" validate:"required,ascii"`
}

// AuthRegister ...
type AuthRegister struct {
	Name     string `json:"name" validate:"required,ascii"`
	Username string `json:"username" validate:"required,ascii"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,ascii"`
}
