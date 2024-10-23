package repository

import (
	"context"
	"errors"
	"time"

	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/google/uuid"
)

// GetDebtByID returns a bebt by ID.
func (repo *Repository) GetDebtByID(ctx context.Context, userID, debtID uuid.UUID) (*domain.Debt, error) {
	query := "SELECT id, created_at, updated_at, deleted_at, user_id, debt_type, amount, description, due_date FROM debts WHERE id = ? AND user_id = ? LIMIT 1"

	bebt := domain.Debt{Base: &domain.Base{}}
	err := repo.db.QueryRowContext(ctx, query, debtID, userID).Scan(
		&bebt.ID,
		&bebt.CreatedAt,
		&bebt.UpdatedAt,
		&bebt.DeletedAt,
		&bebt.UserID,
		&bebt.DebtType,
		&bebt.Amount,
		&bebt.Description,
		&bebt.DueDate,
	)
	if err != nil {
		return nil, err
	}
	return &bebt, nil
}

// GetDebtsByUserID returns the debts of a user by user ID.
func (repo *Repository) GetDebtsByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Debt, error) {
	query := "SELECT id, created_at, updated_at, deleted_at, user_id, debt_type, amount, description, due_date FROM debts WHERE user_id = ?"

	rows, err := repo.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var debts []*domain.Debt
	for rows.Next() {
		debt := domain.Debt{Base: &domain.Base{}}
		if err := rows.Scan(
			&debt.ID,
			&debt.CreatedAt,
			&debt.UpdatedAt,
			&debt.DeletedAt,
			&debt.UserID,
			&debt.DebtType,
			&debt.Amount,
			&debt.Description,
			&debt.DueDate,
		); err != nil {
			return nil, err
		}
		debts = append(debts, &debt)
	}
	return debts, nil
}

// CreateDebt creates a debt.
func (repo *Repository) CreateDebt(ctx context.Context, debt *domain.Debt) (*domain.Debt, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := "INSERT INTO debts (id, created_at, updated_at, user_id, debt_type, amount, description, due_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, query, debt.ID, debt.CreatedAt, debt.UpdatedAt, debt.UserID, debt.DebtType, debt.Amount, debt.Description, debt.DueDate)
	if err != nil {
		return nil, err
	}
	return debt, nil
}

// UpdateDebt updates a debt.
func (repo *Repository) UpdateDebt(ctx context.Context, userID uuid.UUID, debtUpdate *dto.UpdateDebt) (*domain.Debt, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	args := make([]any, 0, 6)
	query := "UPDATE debts SET "
	if debtUpdate.DebtType != nil {
		query += "debt_type = ?, "
		args = append(args, debtUpdate.DebtType)
	}
	if debtUpdate.Amount != nil {
		query += "amount = ?, "
		args = append(args, debtUpdate.Amount)
	}
	if debtUpdate.Description != nil {
		query += "description = ?, "
		args = append(args, debtUpdate.Description)
	}
	if debtUpdate.DueDate != nil {
		query += "due_date = ?, "
		args = append(args, debtUpdate.DueDate)
	}
	query = query[:len(query)-2]

	query += " WHERE id = ? AND user_id = ?"
	args = append(args, debtUpdate.ID, userID)

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	if result, err := result.RowsAffected(); err != nil || result == 0 {
		return nil, errors.New("Debt not found")
	}

	query = "SELECT id, created_at, updated_at, deleted_at, user_id, debt_type, amount, description, due_date FROM debts WHERE user_id = ? LIMIT 1"

	debt := &domain.Debt{Base: &domain.Base{}}
	if err := tx.QueryRowContext(ctx, query, debtUpdate.ID, userID).Scan(
		&debt.ID,
		&debt.CreatedAt,
		&debt.UpdatedAt,
		&debt.DeletedAt,
		&debt.UserID,
		&debt.DebtType,
		&debt.Amount,
		&debt.Description,
		&debt.DueDate,
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return debt, nil
}

// DeleteDebt deletes a debt by debt ID.
func (repo *Repository) DeleteDebt(ctx context.Context, userID uuid.UUID, debtID uuid.UUID) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := "UPDATE debts SET delete_at = ? WHERE id = ? AND user_id = ?"
	result, err := tx.ExecContext(ctx, query, time.Now(), debtID, userID)
	if err != nil {
		return err
	}

	if result, err := result.RowsAffected(); err != nil || result == 0 {
		return errors.New("debt not found")
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
