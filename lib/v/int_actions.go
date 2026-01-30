package v

import (
	"fmt"
)

// numberAction implements IntPipeAction for int validation.
type numberAction struct {
	errorMsg func(v int) string
	validate func(v int) bool
}

// Run executes the validation function on the given int value.
// Returns an error if validation fails.
func (action *numberAction) Run(value int) error {
	if !action.validate(value) {
		return fmt.Errorf("%s", action.errorMsg(value))
	}
	return nil
}

// CustomNumber creates a custom validator using the provided validation function.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	CustomNumber(func(v int) bool { return v%2 == 0 }, ErrMsg{msg: "must be even"})
func CustomNumber(fn func(value int) bool, option ...ActionOptionFace) IntPipeAction {
	return &numberAction{
		errorMsg: func(v int) string {
			return extractMsg("invalid number", v, option...)
		},
		validate: fn,
	}
}

// Gt validates that an int value is strictly greater than the specified value.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	Gt(5) // validates v > 5
func Gt(value int, option ...ActionOptionFace) IntPipeAction {
	return &numberAction{
		errorMsg: func(v int) string {
			return extractMsg("value must be greater than specified value", v, option...)
		},
		validate: func(v int) bool {
			return v > value
		},
	}
}

// Gte validates that an int value is greater than or equal to the specified value.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	Gte(5) // validates v >= 5
func Gte(value int, option ...ActionOptionFace) IntPipeAction {
	return &numberAction{
		errorMsg: func(v int) string {
			return extractMsg("value must be greater than or equal to specified value", v, option...)
		},
		validate: func(v int) bool {
			return v >= value
		},
	}
}

// IsIntString validates that a value can be represented as a valid integer.
// This is a placeholder validator that always returns true.
// The optional ActionOptions parameter can be used to customize the error message.
func IsIntString(option ...ActionOptionFace) IntPipeAction {
	return &numberAction{
		errorMsg: func(v int) string {
			return extractMsg("value must be a valid integer", v, option...)
		},
		validate: func(v int) bool {
			return true
		},
	}
}

// IsNegative validates that an int value is strictly less than zero.
// The optional ActionOptions parameter can be used to customize the error message.
func IsNegative(option ...ActionOptionFace) IntPipeAction {
	return &numberAction{
		errorMsg: func(v int) string {
			return extractMsg("value must be negative", v, option...)
		},
		validate: func(v int) bool {
			return v < 0
		},
	}
}

// IsPositive validates that an int value is strictly greater than zero.
// The optional ActionOptions parameter can be used to customize the error message.
func IsPositive(option ...ActionOptionFace) IntPipeAction {
	return &numberAction{
		errorMsg: func(v int) string {
			return extractMsg("value must be positive", v, option...)
		},
		validate: func(v int) bool {
			return v > 0
		},
	}
}

// IsZero validates that an int value is equal to zero.
// The optional ActionOptions parameter can be used to customize the error message.
func IsZero(option ...ActionOptionFace) IntPipeAction {
	return &numberAction{
		errorMsg: func(v int) string {
			return extractMsg("value must be zero", v, option...)
		},
		validate: func(v int) bool {
			return v == 0
		},
	}
}

// Lt validates that an int value is strictly less than the specified value.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	Lt(10) // validates v < 10
func Lt(value int, option ...ActionOptionFace) IntPipeAction {
	return &numberAction{
		errorMsg: func(v int) string {
			return extractMsg("value must be less than specified value", v, option...)
		},
		validate: func(v int) bool {
			return v < value
		},
	}
}

// Lte validates that an int value is less than or equal to the specified value.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	Lte(10) // validates v <= 10
func Lte(value int, option ...ActionOptionFace) IntPipeAction {
	return &numberAction{
		errorMsg: func(v int) string {
			return extractMsg("value must be less than or equal to specified value", v, option...)
		},
		validate: func(v int) bool {
			return v <= value
		},
	}
}

// Max validates that an int value is less than or equal to the specified maximum.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	Max(100) // validates v <= 100
func Max(max int, option ...ActionOptionFace) IntPipeAction {
	return &numberAction{
		errorMsg: func(v int) string {
			return extractMsg("value exceeds maximum", v, option...)
		},
		validate: func(v int) bool {
			return v <= max
		},
	}
}

// Min validates that an int value is greater than or equal to the specified minimum.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	Min(0) // validates v >= 0
//	Min(10, ErrMsg{msg: "must be at least 10"})
func Min(min int, option ...ActionOptionFace) IntPipeAction {
	return &numberAction{
		errorMsg: func(v int) string {
			return extractMsg("value must be at least specified minimum", v, option...)
		},
		validate: func(v int) bool {
			return v >= min
		},
	}
}
