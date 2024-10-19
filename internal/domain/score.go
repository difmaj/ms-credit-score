package domain

import "time"

// Score represents the credit score of a user
type Score struct {
	*Base
	UserExtended
	Score          int
	LastCalculated time.Time
}
