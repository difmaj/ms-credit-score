package dto

import (
	"time"

	"github.com/difmaj/ms-credit-score/internal/dto/enum"
	"github.com/google/uuid"
)

// DebtOutput represents a user's debt output.
type DebtOutput struct {
	ID          uuid.UUID     `json:"id"`
	UserID      uuid.UUID     `json:"user_id"`
	DebtType    enum.DebtType `json:"debt_type"`
	Amount      float64       `json:"amount"`
	Description string        `json:"description"`
	DueDate     time.Time     `json:"due_date"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeletedAt   *time.Time    `json:"deleted_at"`
}

// GetDebtByIDInput represents a user's debt by ID input.
type GetDebtByIDInput struct {
	DebtID string `uri:"debt_id" validate:"required,uuid"`
}

// GetDebtByIDOutput represents a user's debt by ID output.
type GetDebtByIDOutput DebtOutput

// GetDebtsByUserIDOutput represents a user's debts by user ID output.
type GetDebtsByUserIDOutput struct {
	Debts []*DebtOutput `json:"debts"`
}

// CreateDebtInput represents a user's debt creation input.
type CreateDebtInput struct {
	DebtType    enum.DebtType `json:"debt_type" validate:"required"`
	Amount      float64       `json:"amount" validate:"required,min=0"`
	Description string        `json:"description" validate:"required,min=1,max=255"`
	DueDate     time.Time     `json:"due_date"`
}

// CreateDebtOutput represents a user's debt creation output.
type CreateDebtOutput DebtOutput

// UpdateDebtInput represents a user's debt update input.
type UpdateDebtInput struct {
	DebtType    *enum.DebtType `json:"debt_type"`
	Amount      *float64       `json:"amount" validate:"omitempty,min=0"`
	Description *string        `json:"description" validate:"omitempty,min=1,max=255"`
	DueDate     *time.Time     `json:"due_date"`
}

// UpdateDebtOutput represents a user's debt update output.
type UpdateDebtOutput DebtOutput

// DeleteDebtInput represents a user's debt deletion input.
type DeleteDebtInput struct {
	DebtID string `uri:"debt_id" validate:"required,uuid"`
}

// UpdateDebt represents a user's debt update
type UpdateDebt struct {
	ID          uuid.UUID
	DebtType    *enum.DebtType
	Amount      *float64
	Description *string
	DueDate     *time.Time
}
