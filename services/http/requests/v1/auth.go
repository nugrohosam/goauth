package v1

// AuthLogin ...
type AuthLogin struct {
	EmailOrUsername string `json:"emailOrUsername" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
}
