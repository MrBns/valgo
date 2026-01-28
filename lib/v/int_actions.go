package v

import (
	"fmt"
)

// numberAction implements IntPipeAction for int64 validation.
type numberAction struct {
	errorMsg string
	validate func(v int64) bool
}

// Run executes the validation function on the given int64 value.
// Returns an error if validation fails.
func (action *numberAction) Run(value int64) error {
	if !action.validate(value) {
		return fmt.Errorf("%s", action.errorMsg)
	}
	return nil
}

// Min validates that an int64 value is greater than or equal to the specified minimum.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	Min(0) // validates v >= 0
//	Min(10, ErrMsg{msg: "must be at least 10"})
func Min(min int64, option ...ActionOptions) IntPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be at least specified minimum", option...),
		validate: func(v int64) bool {
			return v >= min
		},
	}
}

// Max validates that an int64 value is less than or equal to the specified maximum.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	Max(100) // validates v <= 100
func Max(max int64, option ...ActionOptions) IntPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value exceeds maximum", option...),
		validate: func(v int64) bool {
			return v <= max
		},
	}
}

// IsPositive validates that an int64 value is strictly greater than zero.
// The optional ActionOptions parameter can be used to customize the error message.
func IsPositive(option ...ActionOptions) IntPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be positive", option...),
		validate: func(v int64) bool {
			return v > 0
		},
	}
}

// IsNegative validates that an int64 value is strictly less than zero.
// The optional ActionOptions parameter can be used to customize the error message.
func IsNegative(option ...ActionOptions) IntPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be negative", option...),
		validate: func(v int64) bool {
			return v < 0
		},
	}
}

// IsZero validates that an int64 value is equal to zero.
// The optional ActionOptions parameter can be used to customize the error message.
func IsZero(option ...ActionOptions) IntPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be zero", option...),
		validate: func(v int64) bool {
			return v == 0
		},
	}
}

// Gt validates that an int64 value is strictly greater than the specified value.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	Gt(5) // validates v > 5
func Gt(value int64, option ...ActionOptions) IntPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be greater than specified value", option...),
		validate: func(v int64) bool {
			return v > value
		},
	}
}

// Gte validates that an int64 value is greater than or equal to the specified value.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	Gte(5) // validates v >= 5
func Gte(value int64, option ...ActionOptions) IntPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be greater than or equal to specified value", option...),
		validate: func(v int64) bool {
			return v >= value
		},
	}
}

// Lt validates that an int64 value is strictly less than the specified value.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	Lt(10) // validates v < 10
func Lt(value int64, option ...ActionOptions) IntPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be less than specified value", option...),
		validate: func(v int64) bool {
			return v < value
		},
	}
}

// Lte validates that an int64 value is less than or equal to the specified value.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	Lte(10) // validates v <= 10
func Lte(value int64, option ...ActionOptions) IntPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be less than or equal to specified value", option...),
		validate: func(v int64) bool {
			return v <= value
		},
	}
}

// IsIntString validates that a value can be represented as a valid integer.
// This is a placeholder validator that always returns true.
// The optional ActionOptions parameter can be used to customize the error message.
func IsIntString(option ...ActionOptions) IntPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be a valid integer", option...),
		validate: func(v int64) bool {
			return true
		},
	}
}

// CustomNumber creates a custom validator using the provided validation function.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	CustomNumber(func(v int64) bool { return v%2 == 0 }, ErrMsg{msg: "must be even"})
func CustomNumber(fn func(value int64) bool, option ...ActionOptions) IntPipeAction {
	return &numberAction{
		errorMsg: extractMsg("invalid number", option...),
		validate: fn,
	}
}
