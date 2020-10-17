package v1

// HeaderJwt using for ...
type HeaderJwt struct {
	Authorization string `header:"Authorization" validate:"required"`
}
