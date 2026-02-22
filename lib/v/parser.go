package v

import (
	"encoding/json"
	"io"
)

// this struct should be embed to the struct
// to make this struct parseable.
//
// This expect to Implement [Schema.Rules] and return a [PipeSet]
// if [Schema.Rules] return nil, it will skip the validation.
type Include struct{}

func (s *Include) Rules() (PipeSet, error) {
	return nil, nil
}

func Validate(s Schema) error {
	rules, err := s.Rules()
	if err != nil {
		return NewPipeError("_pre-check", err)
	}
	if rules == nil {
		return nil
	}
	return rules.Validate()
}

func ValidateAll(s Schema) error {
	rules, err := s.Rules()

	if err != nil {
		return ValidationErrors{NewPipeError("_pre-check", err)}
	}

	if rules == nil {
		return nil
	}

	return rules.ValidateAll()
}

func ValidateAllParallel(s Schema) error {
	return ValidateAll(s)
}

// Parse a schema from [io.Reader] and Validate.
// but if [Schema.Rules] return nil it will skip the validation.
//
// Parse will return only one error which occur first.
//
// Returns nil if there is no error. but return [ParseError] if there is
// any kind of error exists.
func Parse(reader io.Reader, to Schema) error {
	return parseWithDecoder(func(v any) error {
		return json.NewDecoder(reader).Decode(v)
	}, to, false)
}

// ParseFull a schema from [io.Reader] and Validate.
// but if [Schema.Rules] return nil it will skip the validation.
//
// Returns nil if there is no error. but return [ParseError] if there is
// any kind of error exists.
func ParseFull(reader io.Reader, to Schema) error {
	return parseWithDecoder(func(v any) error {
		return json.NewDecoder(reader).Decode(v)
	}, to, true)
}

// ParseBytes a schema from []bytes and Validate.
// but if [Schema.Rules] return nil it will skip the validation.
//
// ParseBytes doesn't return full list of errors, instead
// when the first error happen it return immediately.
func ParseBytes(data []byte, to Schema) error {
	return parseWithDecoder(func(v any) error {
		return json.Unmarshal(data, v)
	}, to, false)
}

// ParseBytesFull a schema from []bytes and Validate.
// but if [Schema.Rules] return nil it will skip the validation.
//
// returns full list of errors.
func ParseBytesFull(data []byte, to Schema) error {
	return parseWithDecoder(func(v any) error {
		return json.Unmarshal(data, v)
	}, to, true)
}

func parseWithDecoder(decode func(any) error, to Schema, full bool) error {
	if err := decode(to); err != nil {
		return &ParseError{ParseError: err}
	}

	pipeSet, err := to.Rules()
	if err != nil {
		return &ParseError{PreError: err}
	}

	if pipeSet == nil {
		return nil
	}

	if full {
		schemaErrors := pipeSet.ValidateAll()
		if schemaErrors == nil {
			return nil
		}
		return &ParseError{ValidationError: schemaErrors}
	}

	schemaError := pipeSet.Validate()
	if schemaError == nil {
		return nil
	}

	return &ParseError{ValidationError: schemaError}
}
