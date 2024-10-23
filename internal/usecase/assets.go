package usecase

import (
	"context"
	"time"

	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/google/uuid"
)

// GetAssetByID returns a asset by ID.
func (u *Usecase) GetAssetByID(ctx context.Context, userID uuid.UUID, in *dto.GetAssetByIDInput) (*dto.AssetOutput, error) {
	result, err := u.repo.GetAssetByID(ctx, userID, uuid.MustParse(in.AssetID))
	if err != nil {
		return nil, err
	}
	return u.assetOutputFromDomain(result), nil
}

// GetAssetsByUserID returns the assets of a user by user ID.
func (u *Usecase) GetAssetsByUserID(ctx context.Context, userID uuid.UUID) ([]*dto.AssetOutput, error) {
	results, err := u.repo.GetAssetsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	out := make([]*dto.AssetOutput, 0, len(results))
	for _, result := range results {
		out = append(out, u.assetOutputFromDomain(result))
	}
	return out, nil
}

// CreateAsset creates a new asset.
func (u *Usecase) CreateAsset(ctx context.Context, userID uuid.UUID, in *dto.CreateAssetInput) (*dto.AssetOutput, error) {
	result, err := u.repo.CreateAsset(ctx, &domain.Asset{
		Base: &domain.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		UserExtended: domain.UserExtended{UserID: userID},
		AssetType:    in.AssetType,
		Value:        in.Value,
		Description:  in.Description,
	})
	if err != nil {
		return nil, err
	}
	return u.assetOutputFromDomain(result), nil
}

// UpdateAsset updates an asset.
func (u *Usecase) UpdateAsset(ctx context.Context, userID uuid.UUID, assertID uuid.UUID, in *dto.UpdateAssetInput) (*dto.AssetOutput, error) {
	result, err := u.repo.UpdateAsset(ctx, userID, &dto.UpdateAsset{
		ID:          assertID,
		AssetType:   in.AssetType,
		Value:       in.Value,
		Description: in.Description,
	})
	if err != nil {
		return nil, err
	}
	return u.assetOutputFromDomain(result), nil
}

// DeleteAsset deletes an asset.
func (u *Usecase) DeleteAsset(ctx context.Context, userID uuid.UUID, in *dto.DeleteAssetInput) error {
	return u.repo.DeleteAsset(ctx, userID, uuid.MustParse(in.AssetID))
}

func (u *Usecase) assetOutputFromDomain(asset *domain.Asset) *dto.AssetOutput {
	return &dto.AssetOutput{
		ID:          asset.ID,
		UserID:      asset.UserID,
		AssetType:   asset.AssetType,
		Value:       asset.Value,
		Description: asset.Description,
		CreatedAt:   asset.CreatedAt,
		UpdatedAt:   asset.UpdatedAt,
		DeletedAt:   asset.DeletedAt,
	}
}
