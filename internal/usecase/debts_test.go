package usecase_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

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

type DebtsSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mock.MockIRepository
	uc   interfaces.IUsecase
}

func (s *DebtsSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock.NewMockIRepository(s.ctrl)
	s.uc = usecase.New(s.repo, nil)
}

func (s *DebtsSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestDebtsSuite(t *testing.T) {
	suite.Run(t, new(DebtsSuite))
}

func (s *DebtsSuite) TestGetDebtByID() {
	s.T().Run("success", func(t *testing.T) {
		userID := uuid.New()
		DebtID := uuid.New()
		input := &dto.GetDebtByIDInput{DebtID: DebtID.String()}

		Debt := &domain.Debt{
			Base: &domain.Base{ID: DebtID},
		}

		s.repo.EXPECT().GetDebtByID(gomock.Any(), userID, DebtID).Return(Debt, nil)

		result, err := s.uc.GetDebtByID(context.Background(), userID, input)
		s.NoError(err)
		s.Equal(DebtID, result.ID)
	})

	s.T().Run("error-debt-not-found", func(t *testing.T) {
		userID := uuid.New()
		DebtID := uuid.New()
		input := &dto.GetDebtByIDInput{DebtID: DebtID.String()}

		s.repo.EXPECT().GetDebtByID(gomock.Any(), userID, DebtID).Return(nil, sql.ErrNoRows)

		result, err := s.uc.GetDebtByID(context.Background(), userID, input)
		s.Error(err)
		s.Nil(result)
	})
}

// Test GetDebtsByUserID
func (s *DebtsSuite) TestGetDebtsByUserID() {
	s.T().Run("success", func(t *testing.T) {
		userID := uuid.New()
		Debts := []*domain.Debt{
			{Base: &domain.Base{ID: uuid.New()}},
			{Base: &domain.Base{ID: uuid.New()}},
		}

		s.repo.EXPECT().GetDebtsByUserID(gomock.Any(), userID).Return(Debts, nil)

		results, err := s.uc.GetDebtsByUserID(context.Background(), userID)
		s.NoError(err)
		s.Len(results, 2)
	})

	s.T().Run("error-no-debts", func(t *testing.T) {
		userID := uuid.New()

		s.repo.EXPECT().GetDebtsByUserID(gomock.Any(), userID).Return(nil, sql.ErrNoRows)

		results, err := s.uc.GetDebtsByUserID(context.Background(), userID)
		s.Error(err)
		s.Nil(results)
	})
}

// Test CreateDebt
func (s *DebtsSuite) TestCreateDebt() {
	s.T().Run("success", func(t *testing.T) {
		userID := uuid.New()
		input := &dto.CreateDebtInput{
			DebtType:    "real estate",
			Amount:      100000.00,
			Description: "Property",
			DueDate:     time.Now().Add(365 * time.Hour * 24),
		}

		Debt := &domain.Debt{
			Base:        &domain.Base{ID: uuid.New()},
			DebtType:    input.DebtType,
			Amount:      input.Amount,
			Description: input.Description,
			DueDate:     input.DueDate,
		}

		s.repo.EXPECT().CreateDebt(gomock.Any(), gomock.Any()).Return(Debt, nil)

		result, err := s.uc.CreateDebt(context.Background(), userID, input)
		s.NoError(err)
		s.NotNil(result)
		s.Equal(enum.DebtType("real estate"), result.DebtType)
	})
}

// Test UpdateDebt
func (s *DebtsSuite) TestUpdateDebt() {
	s.T().Run("success", func(t *testing.T) {
		userID := uuid.New()
		DebtID := uuid.New()
		assertType := enum.DebtTypeMortgage
		dueDate := time.Now().Add(30 * time.Hour * 24)

		input := &dto.UpdateDebtInput{
			DebtType:    &assertType,
			Amount:      pointy.Float64(25000.00),
			Description: pointy.String("Test"),
			DueDate:     &dueDate,
		}

		updatedDebt := &domain.Debt{
			Base:        &domain.Base{ID: DebtID},
			DebtType:    *input.DebtType,
			Amount:      *input.Amount,
			Description: *input.Description,
			DueDate:     *input.DueDate,
		}

		s.repo.EXPECT().UpdateDebt(gomock.Any(), userID, gomock.Any()).Return(updatedDebt, nil)

		result, err := s.uc.UpdateDebt(context.Background(), userID, DebtID, input)
		s.NoError(err)
		s.Equal(enum.DebtTypeMortgage, result.DebtType)
	})
}

// Test DeleteDebt
func (s *DebtsSuite) TestDeleteDebt() {
	s.T().Run("success", func(t *testing.T) {
		userID := uuid.New()
		DebtID := uuid.New()
		input := &dto.DeleteDebtInput{DebtID: DebtID.String()}

		s.repo.EXPECT().DeleteDebt(gomock.Any(), userID, DebtID).Return(nil)

		err := s.uc.DeleteDebt(context.Background(), userID, input)
		s.NoError(err)
	})

	s.T().Run("error-debt-not-found", func(t *testing.T) {
		userID := uuid.New()
		DebtID := uuid.New()
		input := &dto.DeleteDebtInput{DebtID: DebtID.String()}

		s.repo.EXPECT().DeleteDebt(gomock.Any(), userID, DebtID).Return(sql.ErrNoRows)

		err := s.uc.DeleteDebt(context.Background(), userID, input)
		s.Error(err)
	})
}
