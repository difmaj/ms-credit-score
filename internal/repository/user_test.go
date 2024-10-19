package repository_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/difmaj/ms-credit-score/internal/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type RepositoryUserSuite struct {
	suite.Suite
	conn *sql.DB
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	repo *repository.Repository
}

func (rs *RepositoryUserSuite) SetupSuite() {
	var err error
	rs.conn, rs.mock, err = sqlmock.New()
	rs.Require().NoError(err)

	rs.repo, _ = repository.New(rs.conn)
	rs.Require().NotNil(rs.repo)
}

func TestRepositoryUserSuite(t *testing.T) {
	suite.Run(t, new(RepositoryUserSuite))
}

func (rs *RepositoryUserSuite) TestGetUserByEmail() {
	email := "john@example.com"
	query := "SELECT BIN_TO_UUID(id), email, name, password_hash, role_id FROM users WHERE email = ? LIMIT 1"

	rs.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password_hash", "role_id"}).AddRow(uuid.New(), "John Doe", email, "password_hash", uuid.New()))

	user, err := rs.repo.GetUserByEmail(context.Background(), email)
	rs.Require().NoError(err)
	rs.Require().NotNil(user)
	rs.Require().Equal(email, user.Email)
}

func (rs *RepositoryUserSuite) TestGetPrivilegesByUserID() {
	userID := uuid.New()
	query := `
        SELECT p.id, p.context, p.action, p.name, p.description
        FROM privileges p
        JOIN roles_privileges rp ON p.id = rp.privilege_id
        JOIN roles r ON rp.role_id = r.id
        JOIN users u ON u.role_id = r.id
        WHERE u.id = ?`

	rs.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "context", "action", "name", "description"}).
			AddRow(userID, "admin", "read", "Read Access", "Access to read admin data"))

	privileges, err := rs.repo.GetPrivilegesByUserID(context.Background(), userID)
	rs.Require().NoError(err)
	rs.Require().Len(privileges, 1)
	rs.Require().Equal("admin", privileges[0].Context)
}
