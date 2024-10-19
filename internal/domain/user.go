package domain

import "github.com/google/uuid"

// User struct.
type User struct {
	*Base
	Name         string
	Email        string
	PasswordHash string
	RoleExtend
}

// UserExtended struct.
type UserExtended struct {
	UserID uuid.UUID
}
