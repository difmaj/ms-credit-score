package usecase_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go.openly.dev/pointy"

	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/difmaj/ms-credit-score/internal/dto/enum"
	"github.com/difmaj/ms-credit-score/internal/interfaces"
	"github.com/difmaj/ms-credit-score/internal/interfaces/mock"
	"github.com/difmaj/ms-credit-score/internal/usecase"
)

type AssetsSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mock.MockIRepository
	uc   interfaces.IUsecase
}

func (s *AssetsSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock.NewMockIRepository(s.ctrl)
	s.uc = usecase.New(s.repo, nil)
}

func (s *AssetsSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestAssetsSuite(t *testing.T) {
	suite.Run(t, new(AssetsSuite))
}

func (s *AssetsSuite) TestGetAssetByID() {
	s.T().Run("success", func(t *testing.T) {
		userID := uuid.New()
		assetID := uuid.New()
		input := &dto.GetAssetByIDInput{AssetID: assetID.String()}

		asset := &domain.Asset{
			Base: &domain.Base{ID: assetID},
		}

		s.repo.EXPECT().GetAssetByID(gomock.Any(), userID, assetID).Return(asset, nil)

		result, err := s.uc.GetAssetByID(context.Background(), userID, input)
		s.NoError(err)
		s.Equal(assetID, result.ID)
	})

	s.T().Run("error-asset-not-found", func(t *testing.T) {
		userID := uuid.New()
		assetID := uuid.New()
		input := &dto.GetAssetByIDInput{AssetID: assetID.String()}

		s.repo.EXPECT().GetAssetByID(gomock.Any(), userID, assetID).Return(nil, sql.ErrNoRows)

		result, err := s.uc.GetAssetByID(context.Background(), userID, input)
		s.Error(err)
		s.Nil(result)
	})
}

// Test GetAssetsByUserID
func (s *AssetsSuite) TestGetAssetsByUserID() {
	s.T().Run("success", func(t *testing.T) {
		userID := uuid.New()
		assets := []*domain.Asset{
			{Base: &domain.Base{ID: uuid.New()}},
			{Base: &domain.Base{ID: uuid.New()}},
		}

		s.repo.EXPECT().GetAssetsByUserID(gomock.Any(), userID).Return(assets, nil)

		results, err := s.uc.GetAssetsByUserID(context.Background(), userID)
		s.NoError(err)
		s.Len(results, 2)
	})

	s.T().Run("error-no-assets", func(t *testing.T) {
		userID := uuid.New()

		s.repo.EXPECT().GetAssetsByUserID(gomock.Any(), userID).Return(nil, sql.ErrNoRows)

		results, err := s.uc.GetAssetsByUserID(context.Background(), userID)
		s.Error(err)
		s.Nil(results)
	})
}

// Test CreateAsset
func (s *AssetsSuite) TestCreateAsset() {
	s.T().Run("success", func(t *testing.T) {
		userID := uuid.New()
		input := &dto.CreateAssetInput{
			AssetType:   "real estate",
			Value:       100000.00,
			Description: "Property",
		}

		asset := &domain.Asset{
			Base:        &domain.Base{ID: uuid.New()},
			AssetType:   input.AssetType,
			Value:       input.Value,
			Description: input.Description,
		}

		s.repo.EXPECT().CreateAsset(gomock.Any(), gomock.Any()).Return(asset, nil)

		result, err := s.uc.CreateAsset(context.Background(), userID, input)
		s.NoError(err)
		s.NotNil(result)
		s.Equal(enum.AssetType("real estate"), result.AssetType)
	})
}

// Test UpdateAsset
func (s *AssetsSuite) TestUpdateAsset() {
	s.T().Run("success", func(t *testing.T) {
		userID := uuid.New()
		assetID := uuid.New()
		assertType := enum.AssetType("vehicle")

		input := &dto.UpdateAssetInput{
			AssetType:   &assertType,
			Value:       pointy.Float64(25000.00),
			Description: pointy.String("Car"),
		}

		updatedAsset := &domain.Asset{
			Base:        &domain.Base{ID: assetID},
			AssetType:   *input.AssetType,
			Value:       *input.Value,
			Description: *input.Description,
		}

		s.repo.EXPECT().UpdateAsset(gomock.Any(), userID, gomock.Any()).Return(updatedAsset, nil)

		result, err := s.uc.UpdateAsset(context.Background(), userID, assetID, input)
		s.NoError(err)
		s.Equal(enum.AssetType("vehicle"), result.AssetType)
	})
}

// Test DeleteAsset
func (s *AssetsSuite) TestDeleteAsset() {
	s.T().Run("success", func(t *testing.T) {
		userID := uuid.New()
		assetID := uuid.New()
		input := &dto.DeleteAssetInput{AssetID: assetID.String()}

		s.repo.EXPECT().DeleteAsset(gomock.Any(), userID, assetID).Return(nil)

		err := s.uc.DeleteAsset(context.Background(), userID, input)
		s.NoError(err)
	})

	s.T().Run("error-asset-not-found", func(t *testing.T) {
		userID := uuid.New()
		assetID := uuid.New()
		input := &dto.DeleteAssetInput{AssetID: assetID.String()}

		s.repo.EXPECT().DeleteAsset(gomock.Any(), userID, assetID).Return(sql.ErrNoRows)

		err := s.uc.DeleteAsset(context.Background(), userID, input)
		s.Error(err)
	})
}
