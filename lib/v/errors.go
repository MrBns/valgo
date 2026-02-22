package v

import (
	"encoding/json"
	"strings"
)

// PipeError represents a validation error for a specific field.
type PipeError struct {
	Key string
	Err error
}

func NewPipeError(key string, err error) *PipeError {

	if err == nil {
		return nil
	}
	return &PipeError{Key: key, Err: err}
}

func (e *PipeError) Error() string {
	if e.Key == "" {
		return e.Err.Error()
	}
	return e.Key + ": " + e.Err.Error()
}

func (e *PipeError) Unwrap() error {
	return e.Err
}

// MarshalJSON ensures the underlying error string is serialized properly.
func (e *PipeError) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"key": e.Key,
		"msg": e.Err.Error(),
	})
}

// ValidationErrors represents multiple validation errors.
type ValidationErrors []*PipeError

func (v ValidationErrors) Error() string {
	if len(v) == 0 {
		return ""
	}
	var b strings.Builder
	for i, err := range v {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(err.Error())
	}
	return b.String()
}

// Unwrap allows standard library errors.Is and errors.As to work with the slice.
func (v ValidationErrors) Unwrap() []error {
	errs := make([]error, len(v))
	for i, err := range v {
		errs[i] = err
	}
	return errs
}

// ParseError wraps errors that occur during the parsing and validation lifecycle.
type ParseError struct {
	PreError        error `json:"-"`
	ParseError      error `json:"-"`
	ValidationError error `json:"-"`
	PostError       error `json:"-"`
}

// MarshalJSON ensures the underlying errors are serialized properly.
func (e *ParseError) MarshalJSON() ([]byte, error) {
	m := make(map[string]any)
	if e.PreError != nil {
		m["pre_error"] = e.PreError.Error()
	}
	if e.ParseError != nil {
		m["parse_error"] = e.ParseError.Error()
	}
	if e.ValidationError != nil {
		m["validation_error"] = e.ValidationError
	}
	if e.PostError != nil {
		m["post_error"] = e.PostError.Error()
	}
	return json.Marshal(m)
}

func (e *ParseError) Error() string {
	if e.ParseError != nil {
		return "v.parse_error: " + e.ParseError.Error()
	}
	if e.PreError != nil {
		return "v.pre_error: " + e.PreError.Error()
	}
	if e.ValidationError != nil {
		return "v.schema_error: " + e.ValidationError.Error()
	}
	if e.PostError != nil {
		return "v.post_error: " + e.PostError.Error()
	}
	return "unknown parse error"
}

func (e *ParseError) Unwrap() error {
	if e.ParseError != nil {
		return e.ParseError
	}
	if e.PreError != nil {
		return e.PreError
	}
	if e.ValidationError != nil {
		return e.ValidationError
	}
	if e.PostError != nil {
		return e.PostError
	}
	return nil
}
