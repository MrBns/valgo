package v

import (
	"bytes"
	"fmt"
)

// bytesAction implements BytesPipeAction for []byte validation.
type bytesAction struct {
	errorMsg func(v []byte) string
	validate func(v []byte) bool
}

// Run executes the validation function on the given []byte value.
// Returns an error if validation fails.
func (action *bytesAction) Run(value []byte) error {
	if !action.validate(value) {
		return fmt.Errorf("%s", action.errorMsg(value))
	}
	return nil
}

// CustomBytes creates a custom []byte validator using the provided validation function.
// The optional ActionOptions parameter can be used to customize the error message.
func CustomBytes(fn func(value []byte) bool, option ...ActionOptionFace) BytesPipeAction {
	return &bytesAction{
		errorMsg: func(v []byte) string {
			return extractMsg("invalid bytes", string(v), option...)
		},
		validate: fn,
	}
}

// BytesNotEmpty validates that []byte is not empty.
// The optional ActionOptions parameter can be used to customize the error message.
func BytesNotEmpty(option ...ActionOptionFace) BytesPipeAction {
	return &bytesAction{
		errorMsg: func(v []byte) string {
			return extractMsg("cannot be empty", string(v), option...)
		},
		validate: func(v []byte) bool {
			return len(v) > 0
		},
	}
}

// MinBytesLength validates that []byte has at least the specified minimum length.
// The optional ActionOptions parameter can be used to customize the error message.
func MinBytesLength(min int, option ...ActionOptionFace) BytesPipeAction {
	return &bytesAction{
		errorMsg: func(v []byte) string {
			return extractMsg("bytes length must be at least specified minimum", string(v), option...)
		},
		validate: func(v []byte) bool {
			return len(v) >= min
		},
	}
}

// MaxBytesLength validates that []byte does not exceed the specified maximum length.
// The optional ActionOptions parameter can be used to customize the error message.
func MaxBytesLength(max int, option ...ActionOptionFace) BytesPipeAction {
	return &bytesAction{
		errorMsg: func(v []byte) string {
			return extractMsg("bytes length exceeds maximum", string(v), option...)
		},
		validate: func(v []byte) bool {
			return len(v) <= max
		},
	}
}

// EqualBytes validates that []byte is exactly equal to the provided byte slice.
// The optional ActionOptions parameter can be used to customize the error message.
func EqualBytes(target []byte, option ...ActionOptionFace) BytesPipeAction {
	return &bytesAction{
		errorMsg: func(v []byte) string {
			return extractMsg("bytes must be equal to target", string(v), option...)
		},
		validate: func(v []byte) bool {
			return bytes.Equal(v, target)
		},
	}
}

// ContainsBytes validates that []byte contains the provided byte slice.
// The optional ActionOptions parameter can be used to customize the error message.
func ContainsBytes(substr []byte, option ...ActionOptionFace) BytesPipeAction {
	return &bytesAction{
		errorMsg: func(v []byte) string {
			return extractMsg("bytes must contain target bytes", string(v), option...)
		},
		validate: func(v []byte) bool {
			return bytes.Contains(v, substr)
		},
	}
}
