package permission

// TableName ...
const TableName = "permissions"

// Permission struct
type Permission struct {
	ID   int
	Name string
}

// Permissions using for many permissions
type Permissions []Permission
