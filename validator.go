package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/go-universal/i18n"
)

// Validator defines an interface for localized validation functionality.
type Validator interface {
	// AddValidation registers a custom validation rule with a custom validation function.
	AddValidation(rule string, f validator.Func)

	// AddTranslation adds a translation message for a validation rule in a specified locale.
	AddTranslation(locale, rule, message string, options ...i18n.PluralOption)

	// Struct validates an entire struct based on its defined validation rules.
	// Returns validation errors in the provided locale, defaulting to the translator locale if empty.
	Struct(locale string, value any) ValidationError

	// StructExpect validates a struct while ignoring specified fields.
	// Returns validation errors in the provided locale, defaulting to the translator locale if empty.
	StructExpect(locale string, value any, fields ...string) ValidationError

	// StructPartial validates only specified fields of a struct.
	// Returns validation errors in the provided locale, defaulting to the translator locale if empty.
	StructPartial(locale string, value any, fields ...string) ValidationError

	// Var validates a single variable against a rule with optional custom messages.
	Var(locale, name string, value any, rules string) ValidationError

	// VarWithValue validates a variable against another using a rule with custom error messages.
	VarWithValue(locale, name string, value any, other any, rules string) ValidationError
}

type i18nValidator struct {
	prefix     string
	translator i18n.Translator
	validator  *validator.Validate
}

// NewValidator creates a new Validator instance with optional configurations.
// Initializes the I18nValidator with the provided translator and base validator, and applies any options.
func NewValidator(validator *validator.Validate, options ...Options) Validator {
	// Initialize the I18nValidator with the provided validator
	v := &i18nValidator{
		translator: nil,
		validator:  validator,
	}

	// Apply any additional options
	for _, opt := range options {
		opt(v)
	}

	// Return the configured validator instance
	return v
}

func (v *i18nValidator) AddValidation(rule string, f validator.Func) {
	rule = strings.TrimSpace(rule)
	if rule == "" {
		return
	}

	v.validator.RegisterValidation(rule, f)
}

func (v *i18nValidator) AddTranslation(locale, rule, message string, options ...i18n.PluralOption) {
	rule = strings.TrimSpace(rule)
	if rule == "" || v.translator == nil {
		return
	}

	if v.prefix == "" {
		v.translator.AddMessage(locale, rule, message, options...)
	} else {
		v.translator.AddMessage(locale, v.prefix+"."+rule, message, options...)
	}
}

func (v *i18nValidator) Struct(locale string, value any) ValidationError {
	return v.parseStructErrors(
		locale,
		value,
		v.validator.Struct(value),
	)
}

func (v *i18nValidator) StructExpect(locale string, value any, fields ...string) ValidationError {
	return v.parseStructErrors(
		locale,
		value,
		v.validator.StructExcept(value, fields...),
	)
}

func (v *i18nValidator) StructPartial(locale string, value any, fields ...string) ValidationError {
	return v.parseStructErrors(
		locale,
		value,
		v.validator.StructPartial(value, fields...),
	)
}

func (v *i18nValidator) Var(locale, name string, value any, rules string) ValidationError {
	return v.parseVariableErrors(
		locale,
		name,
		value,
		v.validator.Var(value, rules),
	)
}

func (v *i18nValidator) VarWithValue(locale, name string, value any, other any, rules string) ValidationError {
	return v.parseVariableErrors(
		locale,
		name,
		value,
		v.validator.VarWithValue(value, other, rules),
	)
}

// translate generates a localized error message based on the provided value, field, and parameters.
func (v *i18nValidator) translate(locale, name, rule, field string, param, value any, count int) string {
	// Return empty string if translator not passed to I18nValidator
	if v.translator == nil {
		return ""
	}

	// Try resolving error translation using the Translatable interface
	if t, ok := value.(Translatable); ok {
		if res := t.TranslateError(locale, rule, field); res != "" {
			return res
		}
	}

	// Use the main translator to generate the message with pluralization support
	// If a prefix is set, prepend it to the rule
	if v.prefix != "" {
		rule = v.prefix + "." + rule
	}

	// Next, attempt to translate the field name using TranslatableField interface
	if t, ok := value.(TranslatableField); ok {
		if n := t.TranslateTitle(locale, field); n != "" {
			name = n
		}
	}

	return v.translator.Plural(locale, rule, count, map[string]any{
		"field": name,
		"param": param,
	})
}

// parseStructErrors processes and translates validation errors
// based on the provided locale and value for struct.
func (v *i18nValidator) parseStructErrors(locale string, value any, err error) ValidationError {
	// Skip nil error
	if err == nil {
		return NewEmptyError()
	}

	// Assert the error as validator.ValidationErrors
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return NewError(err)
	}

	// Initialize the result validation error
	res := NewEmptyError()

	// Iterate over each validation error and process
	for _, field := range errs {
		// Parse the field parameter and handle its type (int or float)
		var count int
		var param any = field.Param()
		if i, f := parseNumeric(field.Param()); i != nil {
			param = *i
			count = int(*i)
		} else if f != nil {
			param = *f
			count = int(*f)
		}

		// Add the translated error if translator available or raw error to the result
		if v.translator == nil {
			res.AddError(field.Field(), field.Tag(), field.Error())
		} else {
			res.AddError(
				field.Field(),
				field.Tag(),
				v.translate(
					locale, field.Field(), field.Tag(),
					field.StructField(), param, value, count,
				),
			)
		}

	}

	// Return the aggregated validation errors
	return res
}

// parseVariableErrors processes and translates validation errors based on the provided locale and value for variable.
func (v *i18nValidator) parseVariableErrors(locale, name string, value any, err error) ValidationError {
	// Skip nil error
	if err == nil {
		return NewEmptyError()
	}

	// Assert the error as validator.ValidationErrors
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return NewError(err)
	}

	// Initialize the result validation error
	res := NewEmptyError()

	// Iterate over each validation error and process
	for _, field := range errs {
		// Parse the field parameter and handle its type (int or float)
		var count int
		var param any = field.Param()
		if i, f := parseNumeric(field.Param()); i != nil {
			param = *i
			count = int(*i)
		} else if f != nil {
			param = *f
			count = int(*f)
		}

		// Add the translated error if translator available or raw error to the result
		if v.translator == nil {
			res.AddError(field.Field(), field.Tag(), field.Error())
		} else {
			res.AddError(
				name,
				field.Tag(),
				v.translate(
					locale, name, field.Tag(), name,
					param, value, count,
				),
			)
		}
	}

	// Return the aggregated validation errors
	return res
}
