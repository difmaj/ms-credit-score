package response_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/difmaj/ms-credit-score/internal/pkg/router/response"
	"github.com/stretchr/testify/suite"
	"go.uber.org/goleak"
)

type ResponseSuite struct {
	suite.Suite
}

func TestResponseSuite(t *testing.T) {
	suite.Run(t, new(ResponseSuite))
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func (s *ResponseSuite) TestWrite() {
	s.T().Run("valid_response", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		data := map[string]string{"message": "success"}
		resp := response.Response[map[string]string]{
			Success: true,
			Return:  data,
		}

		response.Write(recorder, http.StatusOK, resp)
		s.Require().Equal(http.StatusOK, recorder.Code)
		s.Require().Equal("application/json; charset=UTF-8", recorder.Header().Get("Content-Type"))

		var resBody response.Response[map[string]string]
		err := json.NewDecoder(recorder.Body).Decode(&resBody)
		s.Require().NoError(err)
		s.Require().True(resBody.Success)
		s.Require().Equal(data, resBody.Return)
	})
}

func (s *ResponseSuite) TestError() {
	s.T().Run("single_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		apiErr := &dto.APIError{
			Status:  http.StatusBadRequest,
			Message: "Invalid input",
		}

		response.Error(recorder, apiErr)
		s.Require().Equal(http.StatusBadRequest, recorder.Code)

		var resBody response.Response[any]
		err := json.NewDecoder(recorder.Body).Decode(&resBody)
		s.Require().NoError(err)
		s.Require().False(resBody.Success)
		s.Require().NotNil(resBody.Errors)
		s.Require().Equal(http.StatusBadRequest, resBody.Errors[0].Status)
		s.Require().Equal("Invalid input", resBody.Errors[0].Message)
	})

	s.T().Run("multiple_errors", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		apiErr1 := &dto.APIError{
			Status:  http.StatusBadRequest,
			Message: "First error",
		}
		apiErr2 := &dto.APIError{
			Status:  http.StatusConflict,
			Message: "Second error",
		}

		response.Error(recorder, apiErr1, apiErr2)
		s.Require().Equal(http.StatusBadRequest, recorder.Code)

		var resBody response.Response[any]
		err := json.NewDecoder(recorder.Body).Decode(&resBody)
		s.Require().NoError(err)
		s.Require().False(resBody.Success)
		s.Require().Equal(2, len(resBody.Errors))
		s.Require().Equal("First error", resBody.Errors[0].Message)
		s.Require().Equal("Second error", resBody.Errors[1].Message)
	})
}

func (s *ResponseSuite) TestOk() {
	s.T().Run("success_response", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		data := map[string]string{"message": "all good"}

		response.Ok(recorder, http.StatusOK, data)
		s.Require().Equal(http.StatusOK, recorder.Code)

		var resBody response.Response[map[string]string]
		err := json.NewDecoder(recorder.Body).Decode(&resBody)
		s.Require().NoError(err)
		s.Require().True(resBody.Success)
		s.Require().Equal(data, resBody.Return)
		s.Require().Nil(resBody.Errors)
	})
}
