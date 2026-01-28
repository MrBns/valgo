package v

import (
	"fmt"
)

// floatAction implements FloatPipeAction for float64 validation.
type floatAction struct {
	errorMsg string
	validate func(v float64) bool
}

// Run executes the validation function on the given float64 value.
// Returns an error if validation fails.
func (action *floatAction) Run(value float64) error {
	if !action.validate(value) {
		return fmt.Errorf("%s", action.errorMsg)
	}
	return nil
}

// MinFloat validates that a float64 value is greater than or equal to the specified minimum.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	MinFloat(0.0) // validates v >= 0.0
//	MinFloat(10.5, ErrMsg{msg: "custom error"})
func MinFloat(min float64, option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value must be at least specified minimum", option...),
		validate: func(v float64) bool {
			return v >= min
		},
	}
}

// MaxFloat validates that a float64 value is less than or equal to the specified maximum.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	MaxFloat(100.0) // validates v <= 100.0
func MaxFloat(max float64, option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value exceeds maximum", option...),
		validate: func(v float64) bool {
			return v <= max
		},
	}
}

// IsPositiveFloat validates that a float64 value is greater than or equal to zero.
// The optional ActionOptions parameter can be used to customize the error message.
func IsPositiveFloat(option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value must be positive", option...),
		validate: func(v float64) bool {
			return v >= 0
		},
	}
}

// IsNegativeFloat validates that a float64 value is less than zero.
// The optional ActionOptions parameter can be used to customize the error message.
func IsNegativeFloat(option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value must be negative", option...),
		validate: func(v float64) bool {
			return v < 0
		},
	}
}

// CustomFloat creates a custom validator using the provided validation function.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	CustomFloat(func(v float64) bool { return v != 0 }, ErrMsg{msg: "value cannot be zero"})
func CustomFloat(fn func(value float64) bool, option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("invalid float", option...),
		validate: fn,
	}
}

// GtFloat validates that a float64 value is strictly greater than the specified value.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	GtFloat(5.0) // validates v > 5.0
func GtFloat(value float64, option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value must be greater than specified value", option...),
		validate: func(v float64) bool {
			return v > value
		},
	}
}

// GteFloat validates that a float64 value is greater than or equal to the specified value.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	GteFloat(5.0) // validates v >= 5.0
func GteFloat(value float64, option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value must be greater than or equal to specified value", option...),
		validate: func(v float64) bool {
			return v >= value
		},
	}
}

// LtFloat validates that a float64 value is strictly less than the specified value.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	LtFloat(10.0) // validates v < 10.0
func LtFloat(value float64, option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value must be less than specified value", option...),
		validate: func(v float64) bool {
			return v < value
		},
	}
}

// LteFloat validates that a float64 value is less than or equal to the specified value.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	LteFloat(10.0) // validates v <= 10.0
func LteFloat(value float64, option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value must be less than or equal to specified value", option...),
		validate: func(v float64) bool {
			return v <= value
		},
	}
}
