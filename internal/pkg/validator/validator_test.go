package validator_test

import (
	"net/http"
	"testing"

	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/difmaj/ms-credit-score/internal/pkg/validator"
	"github.com/stretchr/testify/suite"
	"go.uber.org/goleak"
)

type ValidatorSuite struct {
	suite.Suite
	validator *validator.DefaultValidator
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func (s *ValidatorSuite) SetupTest() {
	s.validator = &validator.DefaultValidator{}
}

func TestValidatorSuite(t *testing.T) {
	suite.Run(t, new(ValidatorSuite))
}

func (s *ValidatorSuite) TestValidateStruct() {
	type TestStruct struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}

	s.T().Run("valid_struct", func(t *testing.T) {
		data := TestStruct{Name: "John Doe", Email: "john@example.com"}
		err := s.validator.ValidateStruct(&data)
		s.Require().NoError(err)
	})

	s.T().Run("invalid_struct", func(t *testing.T) {
		data := TestStruct{Name: "", Email: "invalid-email"}
		err := s.validator.ValidateStruct(&data)
		s.Require().Error(err)

		validationErrors, ok := err.(dto.APIErrors)
		s.Require().True(ok)
		s.Require().Equal(2, len(validationErrors))
		s.Require().Equal(http.StatusPreconditionFailed, validationErrors[0].Status)
	})
}
