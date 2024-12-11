package validation

import (
	"errors"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func registerRegexpPattern(v *validator.Validate, tag, pattern string) error {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	return v.RegisterValidation(
		tag,
		func(fl validator.FieldLevel) bool {
			return r.MatchString(fl.Field().String())
		},
	)
}

func SetupCustomValidationRules() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := registerRegexpPattern(v, "retailer", `^[\w\s\-\&]+$`); err != nil {
			return err
		}

		if err := registerRegexpPattern(v, "dateString", `^\d{4}-\d{2}-\d{2}`); err != nil {
			return err
		}

		if err := registerRegexpPattern(v, "timeString", `^\d{2}:\d{2}`); err != nil {
			return err
		}

		if err := registerRegexpPattern(v, "currencyString", `^\d+\.\d{2}$`); err != nil {
			return err
		}

		if err := registerRegexpPattern(v, "shortDescription", `^[\w\s\-]+$`); err != nil {
			return err
		}

		return nil
	} else {
		return errors.New("validation is not available")
	}
}
