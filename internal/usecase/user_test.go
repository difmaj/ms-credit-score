package usecase_test

import (
	"context"
	"testing"

	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/difmaj/ms-credit-score/internal/interfaces"
	"github.com/difmaj/ms-credit-score/internal/interfaces/mock"
	"github.com/difmaj/ms-credit-score/internal/pkg/config"
	"github.com/difmaj/ms-credit-score/internal/usecase"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type UserSuite struct {
	suite.Suite
	ctrl  *gomock.Controller
	repo  *mock.MockIRepository
	redis *mock.MockIRedisClient
	uc    interfaces.IUsecase
}

func (s *UserSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock.NewMockIRepository(s.ctrl)
	s.redis = mock.NewMockIRedisClient(s.ctrl)
	s.uc = usecase.New(s.repo, s.redis)

	err := config.Load("../../.env")
	s.Require().NoError(err)
}

func (s *UserSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (s *UserSuite) TestLogin() {
	s.T().Run("error-invalid-password", func(t *testing.T) {

		email := "john@example.com"
		input := &dto.LoginInput{
			Email:    email,
			Password: "wrong-password",
		}

		user := &domain.User{
			Base:         &domain.Base{ID: uuid.New()},
			Email:        email,
			PasswordHash: "$2a$10$KIXQ8ZKZqgPfhZq.oxcH9e1o1UQk7Ai0a/4qod0x4D4xKPiEKR5oW",
		}

		s.repo.EXPECT().
			GetUserByEmail(gomock.Any(), email).
			Return(user, nil)

		output, err := s.uc.Login(context.Background(), input)
		s.Error(err)
		s.Nil(output)
	})
}
