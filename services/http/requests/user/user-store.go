package user

// UserStoreDto is use
type UserStoreDto struct {
	Name string `json:"name" xml:"name" binding:"required"`
}
