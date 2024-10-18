package domain

// User struct.
type User struct {
	*Base
	Name         string
	Email        string
	PasswordHash string
	RoleExtend
}
