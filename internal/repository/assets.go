package repository

import (
	"context"
	"errors"
	"time"

	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/google/uuid"
)

// GetAssetByID returns a asset by ID.
func (repo *Repository) GetAssetByID(ctx context.Context, userID, assetID uuid.UUID) (*domain.Asset, error) {
	query := "SELECT id, created_at, updated_at, deleted_at, user_id, asset_type, value, description FROM assets WHERE id = ? AND user_id = ? LIMIT 1"

	asset := domain.Asset{Base: &domain.Base{}}
	err := repo.db.QueryRowContext(ctx, query, assetID, userID).Scan(
		&asset.ID,
		&asset.CreatedAt,
		&asset.UpdatedAt,
		&asset.DeletedAt,
		&asset.UserID,
		&asset.AssetType,
		&asset.Value,
		&asset.Description,
	)
	if err != nil {
		return nil, err
	}
	return &asset, nil
}

// GetAssetsByUserID returns the assets of a user by user ID.
func (repo *Repository) GetAssetsByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Asset, error) {
	query := "SELECT id, created_at, updated_at, deleted_at, user_id, asset_type, value, description FROM assets WHERE user_id = ?"

	rows, err := repo.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assets []*domain.Asset
	for rows.Next() {
		asset := domain.Asset{Base: &domain.Base{}}
		if err := rows.Scan(
			&asset.ID,
			&asset.CreatedAt,
			&asset.UpdatedAt,
			&asset.DeletedAt,
			&asset.UserID,
			&asset.AssetType,
			&asset.Value,
			&asset.Description,
		); err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}
	return assets, nil
}

// CreateAsset creates a asset.
func (repo *Repository) CreateAsset(ctx context.Context, asset *domain.Asset) (*domain.Asset, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := "INSERT INTO assets (id, created_at, updated_at, user_id, asset_type, value, description) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, query, asset.ID, asset.CreatedAt, asset.UpdatedAt, asset.UserID, asset.AssetType, asset.Value, asset.Description)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

// UpdateAsset updates a asset.
func (repo *Repository) UpdateAsset(ctx context.Context, userID uuid.UUID, assetUpdate *dto.UpdateAsset) (*domain.Asset, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	args := make([]any, 0, 5)
	query := "UPDATE assets SET "
	if assetUpdate.AssetType != nil {
		query += "asset_type = ?, "
		args = append(args, assetUpdate.AssetType)
	}
	if assetUpdate.Value != nil {
		query += "value = ?, "
		args = append(args, assetUpdate.Value)
	}
	if assetUpdate.Description != nil {
		query += "description = ?, "
		args = append(args, assetUpdate.Description)
	}
	query = query[:len(query)-2]

	query += " WHERE id = ? AND user_id = ?"
	args = append(args, assetUpdate.ID, userID)

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	if result, err := result.RowsAffected(); err != nil || result == 0 {
		return nil, errors.New("asset not found")
	}

	query = "SELECT id, created_at, updated_at, deleted_at, user_id, asset_type, value, description FROM assets WHERE user_id = ? LIMIT 1"

	asset := &domain.Asset{Base: &domain.Base{}}
	if err := tx.QueryRowContext(ctx, query, assetUpdate.ID, userID).Scan(
		&asset.ID,
		&asset.CreatedAt,
		&asset.UpdatedAt,
		&asset.DeletedAt,
		&asset.UserID,
		&asset.AssetType,
		&asset.Value,
		&asset.Description,
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return asset, nil
}

// DeleteAsset deletes a Asset by Asset ID.
func (repo *Repository) DeleteAsset(ctx context.Context, userID uuid.UUID, assetID uuid.UUID) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := "UPDATE assets SET delete_at = ? WHERE id = ? AND user_id = ?"
	result, err := tx.ExecContext(ctx, query, time.Now(), assetID, userID)
	if err != nil {
		return err
	}

	if result, err := result.RowsAffected(); err != nil || result == 0 {
		return errors.New("asset not found")
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
