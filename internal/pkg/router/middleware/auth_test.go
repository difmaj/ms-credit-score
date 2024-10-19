package middleware_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/difmaj/ms-credit-score/internal/pkg/router/middleware"
	"github.com/difmaj/ms-credit-score/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type AuthSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	middleware *middleware.Middleware
	usecase    *mocks.MockIUsecase
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthSuite))
}

func (s *AuthSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.usecase = mocks.NewMockIUsecase(s.ctrl)
	s.middleware = middleware.NewMiddleware(s.usecase)
}

func (s *AuthSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *AuthSuite) TestBasicAuth() {
	s.T().Run("success", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)

		claims := &domain.Claims{
			User:        domain.User{Base: &domain.Base{ID: uuid.New()}},
			Permissions: map[string][]string{},
			Token:       "Bearer token",
			Audience:    os.Getenv("JWT_AUD"),
			Issuer:      os.Getenv("JWT_ISS"),
			Subject:     uuid.New().String(),
			ExpiresAt:   time.Now().Add(1 * time.Hour).Unix(),
		}

		ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
		ctx.Request.Header.Set("Authorization", "Bearer token")

		s.usecase.EXPECT().
			ClaimsJWT("Bearer token").
			Return(claims, nil)

		handler := s.middleware.BasicAuth()
		handler(ctx)

		ctx.Next()

		s.Equal(http.StatusOK, recorder.Code)
		s.Equal(claims.User.ID, ctx.Value("user").(*domain.Claims).User.ID)
	})

	s.T().Run("error-missing-token", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)

		ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)

		handler := s.middleware.BasicAuth()
		handler(ctx)

		ctx.Next()

		s.Equal(http.StatusUnauthorized, recorder.Code)
	})

	s.T().Run("error-invalid-token", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)

		ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
		ctx.Request.Header.Set("Authorization", "Bearer token")

		s.usecase.EXPECT().
			ClaimsJWT("Bearer token").
			Return(nil, errors.New("invalid token"))

		handler := s.middleware.BasicAuth()
		handler(ctx)

		ctx.Next()

		s.Equal(http.StatusUnauthorized, recorder.Code)
	})
}

func (s *AuthSuite) TestPermissionAuth() {
	s.T().Run("success", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)

		claims := &domain.Claims{
			User:        domain.User{Base: &domain.Base{ID: uuid.New()}},
			Permissions: map[string][]string{"context": {"action"}},
			Token:       "Bearer token",
			Audience:    os.Getenv("JWT_AUD"),
			Issuer:      os.Getenv("JWT_ISS"),
			Subject:     uuid.New().String(),
			ExpiresAt:   time.Now().Add(1 * time.Hour).Unix(),
		}

		ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
		ctx.Request.Header.Set("Authorization", "Bearer token")
		ctx.Set("user", claims)

		handler := s.middleware.PermissionAuth("context", "action")
		handler(ctx)

		ctx.Next()

		s.Equal(http.StatusOK, recorder.Code)
	})

	s.T().Run("error-no-permission", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)

		claims := &domain.Claims{
			User:        domain.User{Base: &domain.Base{ID: uuid.New()}},
			Permissions: map[string][]string{"context1": {"action2"}},
		}
		ctx.Set("user", claims)

		ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
		ctx.Request.Header.Set("Authorization", "Bearer token")

		handler := s.middleware.PermissionAuth("context1", "action1")
		handler(ctx)

		ctx.Next()

		s.Equal(http.StatusUnauthorized, recorder.Code)
	})
}
