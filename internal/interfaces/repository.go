package interfaces

import (
	"context"

	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/google/uuid"
)

// IRepository interface.
type IRepository interface {
	// GetUserByID returns a user by ID.
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	// GetPrivilegesByUserID returns the privileges of a user by user ID.
	GetPrivilegesByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Privilege, error)

	// GetAssetByID returns a asset by ID.
	GetAssetByID(ctx context.Context, userID, assetID uuid.UUID) (*domain.Asset, error)
	// GetAssetsByUserID returns the assets of a user by user ID.
	GetAssetsByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Asset, error)
	// CreateAsset creates a new asset.
	CreateAsset(ctx context.Context, asset *domain.Asset) (*domain.Asset, error)
	// UpdateAsset updates an asset.
	UpdateAsset(ctx context.Context, userID uuid.UUID, assetUpdate *dto.UpdateAsset) (*domain.Asset, error)
	// DeleteAsset deletes an asset.
	DeleteAsset(ctx context.Context, userID uuid.UUID, assetID uuid.UUID) error

	// GetDebtByID returns a debt by ID.
	GetDebtByID(ctx context.Context, userID, debtID uuid.UUID) (*domain.Debt, error)
	// GetDebtsByUserID returns the debts of a user by user ID.
	GetDebtsByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Debt, error)
	// CreateDebt creates a new debt.
	CreateDebt(ctx context.Context, debt *domain.Debt) (*domain.Debt, error)
	// UpdateDebt updates a debt.
	UpdateDebt(ctx context.Context, userID uuid.UUID, debtUpdate *dto.UpdateDebt) (*domain.Debt, error)
	// DeleteDebt deletes a debt.
	DeleteDebt(ctx context.Context, userID uuid.UUID, debtID uuid.UUID) error
}
