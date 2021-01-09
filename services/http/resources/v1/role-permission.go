package v1

// RolePermissionItem ...
type RolePermissionItem struct {
	ID           int `structs:"id" json:"id" copier:"field:ID"`
	PermissionID int `structs:"permission_id" json:"permission_id" copier:"field:PermissionID"`
	RoleID       int `structs:"role_id" json:"role_id" copier:"field:RoleID"`
}

// RolePermissionDetail ...
type RolePermissionDetail struct {
	ID           int              `structs:"id" json:"id" copier:"field:ID"`
	Permission   PermissionDetail `structs:"permission" json:"permission" copier:"field:Permission"`
	Role         RoleDetail       `structs:"role" json:"role" copier:"field:Role"`
	PermissionID int              `structs:"permission_id" json:"permission_id" copier:"field:PermissionID"`
	RoleID       int              `structs:"role_id" json:"role_id" copier:"field:RoleID"`
}

// RolePermissionListItems ..
type RolePermissionListItems []RolePermissionItem
