package dto

import (
	"time"

	"github.com/difmaj/ms-credit-score/internal/dto/enum"
	"github.com/google/uuid"
)

// AssetOutput represents a user's asset output.
type AssetOutput struct {
	ID          uuid.UUID      `json:"id"`
	UserID      uuid.UUID      `json:"user_id"`
	AssetType   enum.AssetType `json:"asset_type"`
	Value       float64        `json:"value"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   *time.Time     `json:"deleted_at"`
}

// GetAssetByIDInput represents a user's asset by ID input.
type GetAssetByIDInput struct {
	AssetID string `uri:"asset_id" validate:"required,uuid"`
}

// GetAssetByIDOutput represents a user's asset by ID output.
type GetAssetByIDOutput AssetOutput

// GetAssetsByUserIDOutput represents a user's assets by user ID output.
type GetAssetsByUserIDOutput struct {
	Assets []*AssetOutput `json:"assets"`
}

// CreateAssetInput represents a user's asset creation input.
type CreateAssetInput struct {
	AssetType   enum.AssetType `json:"asset_type" validate:"required"`
	Value       float64        `json:"value" validate:"required,min=0"`
	Description string         `json:"description" validate:"required,min=1,max=255"`
}

// CreateAssetOutput represents a user's asset creation output.
type CreateAssetOutput AssetOutput

// UpdateAssetInput represents a user's asset update input.
type UpdateAssetInput struct {
	AssetType   *enum.AssetType `json:"asset_type"`
	Value       *float64        `json:"value" validate:"omitempty,min=0"`
	Description *string         `json:"description" validate:"omitempty,min=1,max=255"`
}

// UpdateAssetOutput represents a user's asset update output.
type UpdateAssetOutput AssetOutput

// DeleteAssetInput represents a user's asset deletion input.
type DeleteAssetInput struct {
	AssetID string `uri:"asset_id" validate:"required,uuid"`
}

// UpdateAsset represents a user's asset update
type UpdateAsset struct {
	ID          uuid.UUID
	AssetType   *enum.AssetType
	Value       *float64
	Description *string
}
