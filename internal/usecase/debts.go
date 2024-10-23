package usecase

import (
	"context"
	"time"

	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/google/uuid"
)

// GetDebtByID returns a debt by ID.
func (u *Usecase) GetDebtByID(ctx context.Context, userID uuid.UUID, in *dto.GetDebtByIDInput) (*dto.DebtOutput, error) {
	result, err := u.repo.GetDebtByID(ctx, userID, uuid.MustParse(in.DebtID))
	if err != nil {
		return nil, err
	}
	return u.debtOutputFromDomain(result), nil
}

// GetDebtsByUserID returns the debts of a user by user ID.
func (u *Usecase) GetDebtsByUserID(ctx context.Context, userID uuid.UUID) ([]*dto.DebtOutput, error) {
	result, err := u.repo.GetDebtsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	out := make([]*dto.DebtOutput, 0, len(result))
	for _, result := range result {
		out = append(out, u.debtOutputFromDomain(result))
	}
	return out, nil
}

// CreateDebt creates a new debt.
func (u *Usecase) CreateDebt(ctx context.Context, userID uuid.UUID, in *dto.CreateDebtInput) (*dto.DebtOutput, error) {
	result, err := u.repo.CreateDebt(ctx, &domain.Debt{
		Base: &domain.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		UserExtended: domain.UserExtended{UserID: userID},
		DebtType:     in.DebtType,
		Amount:       in.Amount,
		Description:  in.Description,
		DueDate:      in.DueDate,
	})
	if err != nil {
		return nil, err
	}
	return u.debtOutputFromDomain(result), nil
}

// UpdateDebt updates an debt.
func (u *Usecase) UpdateDebt(ctx context.Context, userID uuid.UUID, assertID uuid.UUID, in *dto.UpdateDebtInput) (*dto.DebtOutput, error) {
	result, err := u.repo.UpdateDebt(ctx, userID, &dto.UpdateDebt{
		ID:          assertID,
		DebtType:    in.DebtType,
		Amount:      in.Amount,
		Description: in.Description,
		DueDate:     in.DueDate,
	})
	if err != nil {
		return nil, err
	}
	return u.debtOutputFromDomain(result), nil
}

// DeleteDebt deletes an debt.
func (u *Usecase) DeleteDebt(ctx context.Context, userID uuid.UUID, in *dto.DeleteDebtInput) error {
	return u.repo.DeleteDebt(ctx, userID, uuid.MustParse(in.DebtID))
}

func (u *Usecase) debtOutputFromDomain(debt *domain.Debt) *dto.DebtOutput {
	return &dto.DebtOutput{
		ID:          debt.ID,
		UserID:      debt.UserID,
		DebtType:    debt.DebtType,
		Amount:      debt.Amount,
		Description: debt.Description,
		CreatedAt:   debt.CreatedAt,
		UpdatedAt:   debt.UpdatedAt,
		DeletedAt:   debt.DeletedAt,
	}
}
