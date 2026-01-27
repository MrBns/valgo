package v

import (
	"fmt"
)

type numberAction struct {
	errorMsg string
	validate func(v int64) bool
}

func (action *numberAction) Run(value int64) error {
	if !action.validate(value) {
		return fmt.Errorf("%s", action.errorMsg)
	}
	return nil
}

// Min/Max validators
func IsMinValue(min int64, option ...ActionOptions) NumberPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be at least specified minimum", option...),
		validate: func(v int64) bool {
			return v >= min
		},
	}
}

func IsMaxValue(max int64, option ...ActionOptions) NumberPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value exceeds maximum", option...),
		validate: func(v int64) bool {
			return v <= max
		},
	}
}

// Sign validators
func IsPositive(option ...ActionOptions) NumberPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be positive", option...),
		validate: func(v int64) bool {
			return v > 0
		},
	}
}

func IsNegative(option ...ActionOptions) NumberPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be negative", option...),
		validate: func(v int64) bool {
			return v < 0
		},
	}
}

func IsZero(option ...ActionOptions) NumberPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be zero", option...),
		validate: func(v int64) bool {
			return v == 0
		},
	}
}

// Comparison validators
func IsGt(value int64, option ...ActionOptions) NumberPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be greater than specified value", option...),
		validate: func(v int64) bool {
			return v > value
		},
	}
}

func IsGte(value int64, option ...ActionOptions) NumberPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be greater than or equal to specified value", option...),
		validate: func(v int64) bool {
			return v >= value
		},
	}
}

func IsLt(value int64, option ...ActionOptions) NumberPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be less than specified value", option...),
		validate: func(v int64) bool {
			return v < value
		},
	}
}

func IsLte(value int64, option ...ActionOptions) NumberPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be less than or equal to specified value", option...),
		validate: func(v int64) bool {
			return v <= value
		},
	}
}

// String to int64 validation
func IsIntString(option ...ActionOptions) NumberPipeAction {
	return &numberAction{
		errorMsg: extractMsg("value must be a valid integer", option...),
		validate: func(v int64) bool {
			return true
		},
	}
}

// Custom number validator
func CustomNumber(fn func(value int64) bool, option ...ActionOptions) NumberPipeAction {
	return &numberAction{
		errorMsg: extractMsg("invalid number", option...),
		validate: fn,
	}
}
