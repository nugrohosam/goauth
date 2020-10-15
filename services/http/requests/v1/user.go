package v1

// UserStoreDto is use
type UserStoreDto struct {
	Name     string `form:"name" json:"name" xml:"name" validate:"required"`
	Email    string `form:"email" json:"email" xml:"email" validate:"required"`
	Password string `form:"password" json:"password" xml:"password" validate:"required"`
}
