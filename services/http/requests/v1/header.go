package v1

// HeaderJwt using for ...
type HeaderJwt struct {
	Authorization string `header:"X-Authorization" binding:"required"`
}
