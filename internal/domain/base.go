package domain

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
