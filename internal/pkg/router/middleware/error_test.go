package middleware_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/difmaj/sword-health-backend-challenge/internal/pkg/dtos"
	"github.com/difmaj/sword-health-backend-challenge/internal/pkg/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type ErrorSuite struct {
	suite.Suite
	middleware middleware.IMiddleware
}

func TestErrorSuite(t *testing.T) {
	suite.Run(t, new(ErrorSuite))
}

// SetupSuite initializes the middleware instance
func (s *ErrorSuite) SetupSuite() {
	s.middleware = middleware.NewMiddleware(nil)
}

func (s *ErrorSuite) TestErrorHandler() {
	s.T().Run("no_error", func(t *testing.T) {
		// Setup
		gin.SetMode(gin.TestMode)
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)

		// Apply middleware
		handler := s.middleware.ErrorHandler()
		handler(ctx)

		// Simulate a successful request
		ctx.Next()

		// Assertions
		s.Require().Equal(http.StatusOK, recorder.Code)
		s.Require().Empty(recorder.Body.String())
	})

	s.T().Run("single_api_error", func(t *testing.T) {
		// Setup
		gin.SetMode(gin.TestMode)
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)

		// Add an API error to the context
		apiErr := &dtos.APIError{
			Status:  http.StatusBadRequest,
			Message: "Bad request error",
		}
		ctx.Errors = append(ctx.Errors, &gin.Error{
			Err:  apiErr,
			Type: gin.ErrorTypePrivate,
		})

		// Apply middleware
		handler := s.middleware.ErrorHandler()
		handler(ctx)

		// Simulate request with error
		ctx.Next()

		// Assertions
		s.Require().Equal(http.StatusBadRequest, recorder.Code)

		// Parse response body
		expectedResponse := `{"success":false,"return":null,"errors":[{"status":400,"message":"Bad request error"}]}`
		s.Require().Equal(expectedResponse, recorder.Body.String())
	})

	s.T().Run("multiple_errors", func(t *testing.T) {
		// Setup
		gin.SetMode(gin.TestMode)
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)

		// Add multiple errors to the context
		apiErr := &dtos.APIError{
			Status:  http.StatusBadRequest,
			Message: "Bad request error",
		}
		otherErr := errors.New("some internal error")
		ctx.Errors = append(ctx.Errors, &gin.Error{
			Err:  apiErr,
			Type: gin.ErrorTypePrivate,
		})
		ctx.Errors = append(ctx.Errors, &gin.Error{
			Err:  otherErr,
			Type: gin.ErrorTypePrivate,
		})

		// Apply middleware
		handler := s.middleware.ErrorHandler()
		handler(ctx)

		// Simulate request with errors
		ctx.Next()

		// Assertions
		s.Require().Equal(http.StatusBadRequest, recorder.Code)

		// Parse response body
		expectedResponse := `{"success":false,"return":null,"errors":[{"status":400,"message":"Bad request error"},{"error":"some internal error","status":500,"message":"some internal error"}]}`
		s.Require().Equal(expectedResponse, recorder.Body.String())
	})

	s.T().Run("non_api_error", func(t *testing.T) {
		// Setup
		gin.SetMode(gin.TestMode)
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)

		// Add a non-API error to the context
		otherErr := errors.New("internal server error")
		ctx.Errors = append(ctx.Errors, &gin.Error{
			Err:  otherErr,
			Type: gin.ErrorTypePrivate,
		})

		// Apply middleware
		handler := s.middleware.ErrorHandler()
		handler(ctx)

		// Simulate request with non-API error
		ctx.Next()

		// Assertions
		s.Require().Equal(http.StatusInternalServerError, recorder.Code)

		// Parse response body
		expectedResponse := `{"success":false,"return":null,"errors":[{"error":"internal server error","status":500,"message":"internal server error"}]}`
		s.Require().Equal(expectedResponse, recorder.Body.String())
	})
}
