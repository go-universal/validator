package validator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	v10 "github.com/go-playground/validator/v10"
	"github.com/go-universal/i18n"
	"github.com/go-universal/validator"
	"golang.org/x/text/language"
)

func TestValidator(t *testing.T) {
	// Initialize the validator
	v := validator.NewValidator(
		v10.New(),
		validator.WithTranslator(
			i18n.NewTranslator("en", language.English),
			"",
		),
	)

	v.AddValidation("is_valid", func(fl v10.FieldLevel) bool {
		return fl.Field().String() == "valid"
	})
	v.AddTranslation("en", "is_valid", "{field} must be valid")

	t.Run("Var", func(t *testing.T) {
		err := v.Var("en", "my_field", "valid", "required,is_valid")
		require.False(t, err.HasInternalError(), "unexpected internal error: %v", err.InternalError())
		assert.False(t, err.HasValidationErrors(), "expected no validation errors, got some")
	})

	t.Run("VarWithValue", func(t *testing.T) {
		err := v.VarWithValue("en", "my_field", "value1", "value2", "eqfield")
		require.False(t, err.HasInternalError(), "unexpected internal error: %v", err.InternalError())
		assert.True(t, err.HasValidationErrors(), "expected validation errors, got none")
	})

	t.Run("Struct", func(t *testing.T) {
		type TestStruct struct {
			Field string `validate:"required,is_valid"`
		}
		ts := TestStruct{Field: "valid"}
		err := v.Struct("en", ts)
		require.False(t, err.HasInternalError(), "unexpected internal error: %v", err.InternalError())
		assert.False(t, err.HasValidationErrors(), "expected no validation errors, got some")
	})

	t.Run("InvalidStruct", func(t *testing.T) {
		type TestStruct struct {
			Field string `validate:"required,is_valid"`
		}
		ts := TestStruct{Field: "invalid"}
		err := v.Struct("en", ts)
		require.False(t, err.HasInternalError(), "unexpected internal error: %v", err.InternalError())
		assert.True(t, err.HasValidationErrors(), "expected validation errors, got none")
	})

	t.Run("StructExcept", func(t *testing.T) {
		type TestStruct struct {
			Field1 string `validate:"required,is_valid"`
			Field2 string `validate:"required"`
		}
		ts := TestStruct{Field1: "valid", Field2: ""}
		err := v.StructExpect("en", ts, "Field2")
		require.False(t, err.HasInternalError(), "unexpected internal error: %v", err.InternalError())
		assert.False(t, err.HasValidationErrors(), "expected no validation errors, got some")
	})

	t.Run("StructPartial", func(t *testing.T) {
		type TestStruct struct {
			Field1 string `validate:"required,is_valid"`
			Field2 string `validate:"required"`
		}
		ts := TestStruct{Field1: "invalid", Field2: ""}
		err := v.StructPartial("en", ts, "Field1")
		require.False(t, err.HasInternalError(), "unexpected internal error: %v", err.InternalError())
		assert.True(t, err.HasValidationErrors(), "expected validation errors, got none")
	})
}
