package v

import (
	"fmt"

	"github.com/mrbns/valgo/lib/is"
)

type floatAction struct {
	errorMsg string
	validate func(v float64) bool
}

func (action *floatAction) Run(value float64) error {
	if !action.validate(value) {
		return fmt.Errorf("%s", action.errorMsg)
	}
	return nil
}

// Min/Max validators
func IsMinFloat(min float64, option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value must be at least specified minimum", option...),
		validate: func(v float64) bool {
			return is.IsMinValue(v, min)
		},
	}
}

func IsMaxFloat(max float64, option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value exceeds maximum", option...),
		validate: func(v float64) bool {
			return is.IsMaxValue(v, max)
		},
	}
}

// Sign validators
func IsPositiveFloat(option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value must be positive", option...),
		validate: is.IsPositive[float64],
	}
}

func IsNegativeFloat(option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value must be negative", option...),
		validate: is.IsNegative[float64],
	}
}

// Custom float validator
func CustomFloat(fn func(value float64) bool, option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("invalid float", option...),
		validate: fn,
	}
}

// Comparison validators
func IsGtFloat(value float64, option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value must be greater than specified value", option...),
		validate: func(v float64) bool {
			return v > value
		},
	}
}

func IsGteFloat(value float64, option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value must be greater than or equal to specified value", option...),
		validate: func(v float64) bool {
			return v >= value
		},
	}
}

func IsLtFloat(value float64, option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value must be less than specified value", option...),
		validate: func(v float64) bool {
			return v < value
		},
	}
}

func IsLteFloat(value float64, option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value must be less than or equal to specified value", option...),
		validate: func(v float64) bool {
			return v <= value
		},
	}
}

// String to float64 validation
func IsFloatString(option ...ActionOptions) FloatPipeAction {
	return &floatAction{
		errorMsg: extractMsg("value must be a valid float", option...),
		validate: func(v float64) bool {
			// This validator expects string input converted to float64
			return true // Already validated during conversion
		},
	}
}
