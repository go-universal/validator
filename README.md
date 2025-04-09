# Validator

![GitHub Tag](https://img.shields.io/github/v/tag/go-universal/validator?sort=semver&label=version)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-universal/validator.svg)](https://pkg.go.dev/github.com/go-universal/validator)
[![License](https://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/go-universal/validator/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-universal/validator)](https://goreportcard.com/report/github.com/go-universal/validator)
![Contributors](https://img.shields.io/github/contributors/go-universal/validator)
![Issues](https://img.shields.io/github/issues/go-universal/validator)

`validator` is a Go package that provides a flexible and extensible validation framework with support for localization. It leverages the `go-playground/validator` package for validation and `go-universal/i18n` for localization.

## Features

- Custom validation rules
- Localized error messages
- Struct and variable validation
- Support for various data types and formats

## Installation

To install `validator`, use the following command:

```sh
go get github.com/go-universal/validator
```

## Usage

### Basic Usage

Here's a basic example of how to use `validator`:

```go
package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
    "github.com/go-universal/i18n"
    "github.com/go-universal/validator"
    "golang.org/x/text/language"
)

func main() {
    // Initialize the validator
    v := validator.NewValidator(
        validator.New(),
        validator.WithTranslator(
            i18n.NewTranslator("en", language.English),
            "",
        ),
    )

    // Add custom validation rule
    v.AddValidation("is_valid", func(fl validator.FieldLevel) bool {
        return fl.Field().String() == "valid"
    })
    v.AddTranslation("en", "is_valid", "{field} must be valid")

    // Validate a variable
    err := v.Var("en", "my_field", "valid", "required,is_valid")
    if err.HasInternalError() {
        fmt.Println("Internal error:", err.InternalError())
    } else if err.HasValidationErrors() {
        fmt.Println("Validation errors:", err)
    } else {
        fmt.Println("Validation passed")
    }
}
```

### Struct Validation

You can also validate structs with `validator`:

```go
package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
    "github.com/go-universal/i18n"
    "github.com/go-universal/validator"
    "golang.org/x/text/language"
)

type TestStruct struct {
    Field string `validate:"required,is_valid"`
}

func main() {
    // Initialize the validator
    v := validator.NewValidator(
        validator.New(),
        validator.WithTranslator(
            i18n.NewTranslator("en", language.English),
            "",
        ),
    )

    // Add custom validation rule
    v.AddValidation("is_valid", func(fl validator.FieldLevel) bool {
        return fl.Field().String() == "valid"
    })
    v.AddTranslation("en", "is_valid", "{field} must be valid")

    // Validate a struct
    ts := TestStruct{Field: "valid"}
    err := v.Struct("en", ts)
    if err.HasInternalError() {
        fmt.Println("Internal error:", err.InternalError())
    } else if err.HasValidationErrors() {
        fmt.Println("Validation errors:", err)
    } else {
        fmt.Println("Validation passed")
    }
}
```

### Custom Validators

You can add custom validators to `validator`:

```go
package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
    "github.com/go-universal/i18n"
    "github.com/go-universal/validator"
    "golang.org/x/text/language"
)

func main() {
    // Initialize the validator
    v := validator.NewValidator(
        validator.New(),
        validator.WithTranslator(
            i18n.NewTranslator("en", language.English),
            "",
        ),
    )

    // Add custom validation rule
    v.AddValidation("is_valid", func(fl validator.FieldLevel) bool {
        return fl.Field().String() == "valid"
    })
    v.AddTranslation("en", "is_valid", "{field} must be valid")

    // Validate a variable
    err := v.Var("en", "my_field", "invalid", "required,is_valid")
    if err.HasInternalError() {
        fmt.Println("Internal error:", err.InternalError())
    } else if err.HasValidationErrors() {
        fmt.Println("Validation errors:", err)
    } else {
        fmt.Println("Validation passed")
    }
}
```

## License

This project is licensed under the ISC License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [go-playground/validator](https://github.com/go-playground/validator)
- [go-universal/i18n](https://github.com/go-universal/i18n)
