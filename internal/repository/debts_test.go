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

type RepositoryDebtSuite struct {
	suite.Suite
	conn *sql.DB
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	repo *repository.Repository
}

func (rs *RepositoryDebtSuite) SetupTest() {
	var err error
	rs.conn, rs.mock, err = sqlmock.New()
	rs.Require().NoError(err)

	rs.repo, _ = repository.New(rs.conn)
	rs.Require().NotNil(rs.repo)
}

func TestRepositoryDebtSuite(t *testing.T) {
	suite.Run(t, new(RepositoryDebtSuite))
}

func (rs *RepositoryDebtSuite) TestGetDebtByID() {
	userID := uuid.New()
	debtID := uuid.New()
	query := "SELECT id, created_at, updated_at, deleted_at, user_id, debt_type, amount, description, due_date FROM debts WHERE id = ? AND user_id = ? LIMIT 1"

	rs.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(debtID, userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "user_id", "debt_type", "amount", "description", "due_date"}).
			AddRow(debtID, time.Now(), time.Now(), nil, userID, "loan", 5000, "Car loan", time.Now().Add(30*time.Hour*24)))

	debt, err := rs.repo.GetDebtByID(context.Background(), userID, debtID)
	rs.Require().NoError(err)
	rs.Require().NotNil(debt)
	rs.Require().Equal(debtID, debt.ID)
	rs.Require().Equal(userID, debt.UserID)
}

func (rs *RepositoryDebtSuite) TestGetDebtsByUserID() {
	userID := uuid.New()
	query := "SELECT id, created_at, updated_at, deleted_at, user_id, debt_type, amount, description, due_date FROM debts WHERE user_id = ?"

	rs.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "user_id", "debt_type", "amount", "description", "due_date"}).
			AddRow(uuid.New(), time.Now(), time.Now(), nil, userID, "credit card", 3000, "Credit card debt", time.Now().Add(30*time.Hour*24)).
			AddRow(uuid.New(), time.Now(), time.Now(), nil, userID, "mortgage", 150000, "Mortgage debt", time.Now().Add(365*time.Hour*24)))

	debts, err := rs.repo.GetDebtsByUserID(context.Background(), userID)
	rs.Require().NoError(err)
	rs.Require().Len(debts, 2)
	rs.Require().Equal(enum.DebtType("credit card"), debts[0].DebtType)
	rs.Require().Equal(enum.DebtType("mortgage"), debts[1].DebtType)
}

func (rs *RepositoryDebtSuite) TestCreateDebt() {
	userID := uuid.New()
	debt := &domain.Debt{
		Base: &domain.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		UserExtended: domain.UserExtended{UserID: userID},
		DebtType:     "personal loan",
		Amount:       10000,
		Description:  "Personal loan for education",
		DueDate:      time.Now().Add(365 * time.Hour * 24),
	}

	rs.mock.ExpectBegin()
	query := "INSERT INTO debts (id, created_at, updated_at, user_id, debt_type, amount, description, due_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	rs.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), debt.UserID, debt.DebtType, debt.Amount, debt.Description, debt.DueDate).
		WillReturnResult(sqlmock.NewResult(1, 1))
	rs.mock.ExpectCommit()

	createdDebt, err := rs.repo.CreateDebt(context.Background(), debt)
	rs.Require().NoError(err)
	rs.Require().NotNil(createdDebt)
	rs.Require().Equal(debt.UserID, createdDebt.UserID)
}

func (rs *RepositoryDebtSuite) TestUpdateDebt() {
	userID := uuid.New()
	debtID := uuid.New()
	updatedDueDate := time.Now().Add(30 * time.Hour * 24)

	debtUpdate := &dto.UpdateDebt{
		ID:          debtID,
		Amount:      pointy.Float64(12000),
		Description: pointy.String("Updated personal loan for education"),
		DueDate:     &updatedDueDate,
	}

	rs.mock.ExpectBegin()
	query := "UPDATE debts SET amount = ?, description = ?, due_date = ? WHERE id = ? AND user_id = ?"
	rs.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(debtUpdate.Amount, debtUpdate.Description, debtUpdate.DueDate, debtUpdate.ID, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	selectQuery := "SELECT id, created_at, updated_at, deleted_at, user_id, debt_type, amount, description, due_date FROM debts WHERE user_id = ? LIMIT 1"
	rs.mock.ExpectQuery(regexp.QuoteMeta(selectQuery)).
		WithArgs(debtID, userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "user_id", "debt_type", "amount", "description", "due_date"}).
			AddRow(debtID, time.Now(), time.Now(), nil, userID, "loan", debtUpdate.Amount, debtUpdate.Description, debtUpdate.DueDate))

	rs.mock.ExpectCommit()

	updatedDebt, err := rs.repo.UpdateDebt(context.Background(), userID, debtUpdate)
	rs.Require().NoError(err)
	rs.Require().NotNil(updatedDebt)
	rs.Require().Equal(*debtUpdate.Amount, updatedDebt.Amount)
	rs.Require().Equal(*debtUpdate.Description, updatedDebt.Description)
	rs.Require().Equal(updatedDueDate, updatedDebt.DueDate)
}

func (rs *RepositoryDebtSuite) TestDeleteDebt() {
	debtID := uuid.New()
	userID := uuid.New()
	query := "UPDATE debts SET delete_at = ? WHERE id = ? AND user_id = ?"

	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(sqlmock.AnyArg(), debtID, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	rs.mock.ExpectCommit()

	err := rs.repo.DeleteDebt(context.Background(), userID, debtID)
	rs.Require().NoError(err)
}
