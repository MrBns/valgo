package v

import (
	"encoding/json"
	"io"
)

// this struct should be embedd to the struct
// to make this struct parseable.
//
// This expect to Implement [Schema.Rules] and return a [Pipeset]
// if [Schema.Rules] return nil, it will skip the validation.
type Include struct{}

func (s *Include) Rules() (Pipeset, error) {
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
				&SchemaError{
					Key: "_pre-check",
					Err: err,
				},
			},
		}
	}

	return rules.ValidateAll()
}

// Parse a schema from [io.Reader] and Validate.
// but if [Schema.Rules] return nil it will skip the validation.
func Parse(reader io.Reader, to Schema) *ParseErrors {

	var parseErrors *ParseErrors = nil

	err := json.NewDecoder(reader).Decode(to)
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

	errors := pipeset.ValidateAll()

	return &ParseErrors{
		SchemaErrors: errors,
	}
}

// Parse a schema from []bytes and Validate.
// but if [Schema.Rules] return nil it will skip the validation.
func ParseBytes(data []byte, to Schema) *ParseErrors {

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
