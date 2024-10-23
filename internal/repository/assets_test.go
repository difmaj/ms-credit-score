package repository_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/difmaj/ms-credit-score/internal/dto/enum"
	"github.com/difmaj/ms-credit-score/internal/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go.openly.dev/pointy"
	"gorm.io/gorm"
)

type RepositoryAssetSuite struct {
	suite.Suite
	conn *sql.DB
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	repo *repository.Repository
}

func (rs *RepositoryAssetSuite) SetupTest() {
	var err error
	rs.conn, rs.mock, err = sqlmock.New()
	rs.Require().NoError(err)

	rs.repo, _ = repository.New(rs.conn)
	rs.Require().NotNil(rs.repo)
}

func TestRepositoryAssetSuite(t *testing.T) {
	suite.Run(t, new(RepositoryAssetSuite))
}

func (rs *RepositoryAssetSuite) TestGetAssetByID() {
	userID := uuid.New()
	assetID := uuid.New()
	query := "SELECT id, created_at, updated_at, deleted_at, user_id, asset_type, value, description FROM assets WHERE id = ? AND user_id = ? LIMIT 1"

	rs.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(assetID, userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "user_id", "asset_type", "value", "description"}).
			AddRow(assetID, time.Now(), time.Now(), nil, userID, "property", 100000, "A house"))

	asset, err := rs.repo.GetAssetByID(context.Background(), userID, assetID)
	rs.Require().NoError(err)
	rs.Require().NotNil(asset)
	rs.Require().Equal(assetID, asset.ID)
	rs.Require().Equal(userID, asset.UserID)
}

func (rs *RepositoryAssetSuite) TestGetAssetsByUserID() {
	userID := uuid.New()
	query := "SELECT id, created_at, updated_at, deleted_at, user_id, asset_type, value, description FROM assets WHERE user_id = ?"

	rs.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "user_id", "asset_type", "value", "description"}).
			AddRow(uuid.New(), time.Now(), time.Now(), nil, userID, "vehicle", 30000, "A car").
			AddRow(uuid.New(), time.Now(), time.Now(), nil, userID, "property", 150000, "An apartment"))

	assets, err := rs.repo.GetAssetsByUserID(context.Background(), userID)
	rs.Require().NoError(err)
	rs.Require().Len(assets, 2)
	rs.Require().Equal(enum.AssetType("vehicle"), assets[0].AssetType)
	rs.Require().Equal(enum.AssetType("property"), assets[1].AssetType)
}

func (rs *RepositoryAssetSuite) TestCreateAsset() {
	userID := uuid.New()
	asset := &domain.Asset{
		Base: &domain.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		UserExtended: domain.UserExtended{
			UserID: userID,
		},
		AssetType:   "stock",
		Value:       50000,
		Description: "Company stocks",
	}

	rs.mock.ExpectBegin()
	query := "INSERT INTO assets (id, created_at, updated_at, user_id, asset_type, value, description) VALUES (?, ?, ?, ?, ?, ?, ?)"
	rs.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), asset.UserID, asset.AssetType, asset.Value, asset.Description).
		WillReturnResult(sqlmock.NewResult(1, 1))
	rs.mock.ExpectCommit()

	createdAsset, err := rs.repo.CreateAsset(context.Background(), asset)
	rs.Require().NoError(err)
	rs.Require().NotNil(createdAsset)
	rs.Require().Equal(asset.UserID, createdAsset.UserID)
}

func (rs *RepositoryAssetSuite) TestUpdateAsset() {
	userID := uuid.New()
	assetID := uuid.New()

	assetUpdate := &dto.UpdateAsset{
		ID:          assetID,
		Value:       pointy.Float64(75000),
		Description: pointy.String("Updated asset description"),
	}

	rs.mock.ExpectBegin()
	query := "UPDATE assets SET value = ?, description = ? WHERE id = ? AND user_id = ?"
	rs.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(assetUpdate.Value, assetUpdate.Description, assetUpdate.ID, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	selectQuery := "SELECT id, created_at, updated_at, deleted_at, user_id, asset_type, value, description FROM assets WHERE user_id = ? LIMIT 1"
	rs.mock.ExpectQuery(regexp.QuoteMeta(selectQuery)).
		WithArgs(assetID, userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "user_id", "asset_type", "value", "description"}).
			AddRow(assetID, time.Now(), time.Now(), nil, userID, "property", assetUpdate.Value, assetUpdate.Description))

	rs.mock.ExpectCommit()

	updatedAsset, err := rs.repo.UpdateAsset(context.Background(), userID, assetUpdate)
	rs.Require().NoError(err)
	rs.Require().NotNil(updatedAsset)
	rs.Require().Equal(*assetUpdate.Value, updatedAsset.Value)
	rs.Require().Equal(*assetUpdate.Description, updatedAsset.Description)
}

func (rs *RepositoryAssetSuite) TestDeleteAsset() {
	assetID := uuid.New()
	userID := uuid.New()
	query := "UPDATE assets SET delete_at = ? WHERE id = ? AND user_id = ?"

	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(sqlmock.AnyArg(), assetID, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	rs.mock.ExpectCommit()

	err := rs.repo.DeleteAsset(context.Background(), userID, assetID)
	rs.Require().NoError(err)
}
