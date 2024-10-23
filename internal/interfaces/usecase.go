package interfaces

import (
	"context"

	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/google/uuid"
)

// IUsecase represents the usecase interface.
type IUsecase interface {
	// ClaimsJWT returns the claims from the JWT token.
	ClaimsJWT(token string) (*domain.Claims, error)

	// Login handles the login request.
	Login(context.Context, *dto.LoginInput) (*dto.LoginOutput, error)

	// GetAssetByID returns an asset by its ID.
	GetAssetByID(ctx context.Context, userID uuid.UUID, in *dto.GetAssetByIDInput) (*dto.AssetOutput, error)
	// GetAssetsByUserID returns the assets of a user by user ID.
	GetAssetsByUserID(ctx context.Context, userID uuid.UUID) ([]*dto.AssetOutput, error)
	// CreateAsset creates a new asset.
	CreateAsset(ctx context.Context, userID uuid.UUID, in *dto.CreateAssetInput) (*dto.AssetOutput, error)
	// UpdateAsset updates an asset.
	UpdateAsset(ctx context.Context, userID uuid.UUID, assertID uuid.UUID, in *dto.UpdateAssetInput) (*dto.AssetOutput, error)
	// DeleteAsset deletes an asset.
	DeleteAsset(ctx context.Context, userID uuid.UUID, in *dto.DeleteAssetInput) error

	// GetDebtByID returns a debt by ID.
	GetDebtByID(ctx context.Context, userID uuid.UUID, in *dto.GetDebtByIDInput) (*dto.DebtOutput, error)
	// GetDebtsByUserID returns the debts of a user by user ID.
	GetDebtsByUserID(ctx context.Context, userID uuid.UUID) ([]*dto.DebtOutput, error)
	// CreateDebt creates a new debt.
	CreateDebt(ctx context.Context, userID uuid.UUID, in *dto.CreateDebtInput) (*dto.DebtOutput, error)
	// UpdateDebt updates an debt.
	UpdateDebt(ctx context.Context, userID uuid.UUID, assertID uuid.UUID, in *dto.UpdateDebtInput) (*dto.DebtOutput, error)
	// DeleteDebt deletes an debt.
	DeleteDebt(ctx context.Context, userID uuid.UUID, in *dto.DeleteDebtInput) error
}
