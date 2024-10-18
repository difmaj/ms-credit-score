package domain

import "github.com/google/uuid"

type Role struct {
	*Base
	Name        string
	Description string
}

type RoleExtend struct {
	RoleID uuid.UUID
}

type Privilege struct {
	*Base
	Name        string
	Description string
	Context     string
	Action      string
}

type PrivilegeExtend struct {
	PrivilegeID uuid.UUID
}

type RolePrivilege struct {
	*Base
	PrivilegeExtend
	RoleExtend
}
