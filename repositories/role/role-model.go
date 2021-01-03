package role

// TableName ...
const TableName = "roles"

// Role struct
type Role struct {
	ID   int
	Name string
}

// Roles using for many roles
type Roles []Role
