package validator

import (
	"net/http"
	"reflect"
	"sync"

	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var tagsToLookup = []string{"json", "form", "uri"}

// Ensure defaultValidator implements binding.StructValidator interface
var _ binding.StructValidator = &DefaultValidator{}

// DefaultValidator implements Gin's binding.StructValidator interface
type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
}

// lazyinit initializes the validator and translation based on the provided language
func (v *DefaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		en := en.New()
		v.uni = ut.New(en, en)

		v.trans, _ = v.uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(v.validate, v.trans)
	})
}

// ValidateStruct validates a struct and returns translated errors if any
func (v *DefaultValidator) ValidateStruct(s interface{}) error {
	if kindOfData(s) == reflect.Struct {
		v.lazyinit()

		if err := v.validate.Struct(s); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				var exs dto.APIErrors
				stype := reflect.TypeOf(s)
				if stype.Kind() == reflect.Ptr {
					stype = stype.Elem()
				}
				for _, err := range errs {
					exception := dto.NewAPIError(http.StatusPreconditionFailed, err, err.Translate(v.trans))
					exs = append(exs, exception)
				}
				return exs
			}
		}
	}
	return nil
}

// Engine returns the underlying validator engine
func (v *DefaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

// kindOfData determines the kind of the given data, following pointers if necessary
func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
