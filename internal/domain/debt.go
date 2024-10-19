package domain

import (
	"time"

	"github.com/difmaj/ms-credit-score/internal/dto/enum"
)

// Debt represents a user's debt
type Debt struct {
	*Base
	UserExtended
	DebtType    enum.DebtType
	Amount      float64
	Description string
	DueDate     time.Time
}
