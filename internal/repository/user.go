package repository

import (
	"context"

	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/google/uuid"
)

// GetUserByEmail returns a user by email.
func (repo *Repository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := domain.User{
		Base: &domain.Base{},
	}
	if err := repo.db.QueryRowContext(ctx, "SELECT BIN_TO_UUID(id), email, name, password_hash, role_id FROM users WHERE email = ? LIMIT 1", email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.RoleID,
	); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetPrivilegesByUserID returns the privileges of a user by user ID.
// Decided to use raw SQL query because of the query performance.
func (repo *Repository) GetPrivilegesByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Privilege, error) {
	query := `
        SELECT p.id, p.context, p.action, p.name, p.description
        FROM privileges p
        JOIN roles_privileges rp ON p.id = rp.privilege_id
        JOIN roles r ON rp.role_id = r.id
        JOIN users u ON u.role_id = r.id
        WHERE u.id = ?`

	rows, err := repo.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var privileges []*domain.Privilege
	for rows.Next() {
		privilege := &domain.Privilege{Base: &domain.Base{}}
		if err := rows.Scan(&privilege.ID, &privilege.Context, &privilege.Action, &privilege.Name, &privilege.Description); err != nil {
			return nil, err
		}
		privileges = append(privileges, privilege)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return privileges, nil
}
