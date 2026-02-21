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

func Validate(s Schema) *SchemaError {
	rules, err := s.Rules()
	if err != nil {
		return &SchemaError{
			Key: "_pre-check",
			Err: err,
		}
	}
	return rules.Validate()
}

func ValidateAll(s Schema) *SchemaErrors {
	rules, err := s.Rules()

	if err != nil {
		return &SchemaErrors{
			Errors: []*SchemaError{
				{
					Key: "_pre-check",
					Err: err,
				},
			},
		}
	}

	return rules.ValidateAll()
}

func ValidateAllParallel(s Schema) *SchemaErrors {
	return ValidateAll(s)
}

// Parse a schema from [io.Reader] and Validate.
// but if [Schema.Rules] return nil it will skip the validation.
//
// Parse will return only one error which occur first.
//
// Returns nil if there is no error. but return [ParseErrors] if there is
// any kind of error exists.
func Parse(reader io.Reader, to Schema) *ParseErrors {

	var parseErrors *ParseErrors = nil

	err := json.NewDecoder(reader).Decode(to)
	if err != nil {
		parseErrors = &ParseErrors{
			ParseError: err,
		}
		return parseErrors
	}

	pipeSet, err := to.Rules()
	if err != nil {
		return &ParseErrors{
			PreErrors: err,
		}
	}

	if pipeSet == nil {
		return nil
	}

	parseErr := pipeSet.Validate()
	if parseErr == nil {
		return nil
	}

	return &ParseErrors{
		SchemaErrors: &SchemaErrors{
			Errors: []*SchemaError{parseErr},
		},
	}
}

// ParseFull a schema from [io.Reader] and Validate.
// but if [Schema.Rules] return nil it will skip the validation.
//
// Returns nil if there is no error. but return [ParseErrors] if there is
// any kind of error exists.
func ParseFull(reader io.Reader, to Schema) *ParseErrors {

	var parseErrors *ParseErrors = nil

	err := json.NewDecoder(reader).Decode(to)
	if err != nil {
		parseErrors = &ParseErrors{
			ParseError: err,
		}
		return parseErrors
	}

	pipeSet, err := to.Rules()
	if err != nil {
		return &ParseErrors{
			PreErrors: err,
		}
	}

	if pipeSet == nil {
		return nil
	}

	errors := pipeSet.ValidateAll()
	if errors != nil && len(errors.Errors) > 0 {

		return &ParseErrors{
			SchemaErrors: errors,
		}
	}
	return nil

}

// ParseBytes a schema from []bytes and Validate.
// but if [Schema.Rules] return nil it will skip the validation.
//
// ParseBytes doesn't return full list of errors, instead
// when the first error happen it return immediately.
func ParseBytes(data []byte, to Schema) *ParseErrors {

	var parseErrors *ParseErrors = nil

	err := json.Unmarshal(data, to)
	if err != nil {
		parseErrors = &ParseErrors{
			ParseError: err,
		}
		return parseErrors
	}

	pipeSet, err := to.Rules()
	if err != nil {
		return &ParseErrors{
			PreErrors: err,
		}
	}

	if pipeSet == nil {
		return nil
	}

	parseErr := pipeSet.Validate()

	return &ParseErrors{
		SchemaErrors: &SchemaErrors{
			Errors: []*SchemaError{
				parseErr,
			},
		},
	}
}

// ParseBytesFull a schema from []bytes and Validate.
// but if [Schema.Rules] return nil it will skip the validation.
//
// returns full list of errors.
func ParseBytesFull(data []byte, to Schema) *ParseErrors {

	var parseErrors *ParseErrors = nil

	err := json.Unmarshal(data, to)
	if err != nil {
		parseErrors = &ParseErrors{
			ParseError: err,
		}
		return parseErrors
	}

	pipeset, err := to.Rules()
	if err != nil {
		return &ParseErrors{
			PreErrors: err,
		}
	}

	if pipeset == nil {
		return nil
	}

	parseErros := pipeset.ValidateAll()

	return &ParseErrors{
		SchemaErrors: parseErros,
	}
}
