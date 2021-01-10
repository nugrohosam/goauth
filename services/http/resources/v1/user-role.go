package v1

// UserRoleItem ...
type UserRoleItem struct {
	ID     int `structs:"id" json:"id" copier:"field:ID"`
	UserID int `structs:"user_id" json:"user_id" copier:"field:UserID"`
	RoleID int `structs:"role_id" json:"role_id" copier:"field:RoleID"`
}

// UserRoleDetail ...
type UserRoleDetail struct {
	ID           int        `structs:"id" json:"id" copier:"field:ID"`
	User         UserDetail `structs:"user" json:"user" copier:"field:User"`
	Role         RoleDetail `structs:"role" json:"role" copier:"field:Role"`
	PermissionID int        `structs:"user_id" json:"user_id" copier:"field:UserID"`
	RoleID       int        `structs:"role_id" json:"role_id" copier:"field:RoleID"`
}

// UserRoleListItems ..
type UserRoleListItems []UserRoleItem
