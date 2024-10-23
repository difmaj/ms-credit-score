package domain

import (
	"github.com/difmaj/ms-credit-score/internal/dto/enum"
)

// Asset represents a user's asset
type Asset struct {
	*Base
	UserExtended
	AssetType   enum.AssetType
	Value       float64
	Description string
}
