package v

import "errors"

// SchemaError represents a validation error with a key and an error message.
type SchemaError struct {
	Key string `json:"key"`
	Err error  `json:"msg"`
}

// SchemaErrors is a slice of SchemaError pointers representing multiple validation errors.
type SchemaErrors struct {
	Errors []*SchemaError `json:"errors"`
}

type ParseErrors struct {
	// Validation Errors.
	*SchemaErrors
	// Errors happened before validation operation.
	PreErrors error `json:"pre_error"`
	// Parse Error. error during json parsing.
	ParseError error `json:"parse_error"`
	// Errors happened after valiation operation.
	// if validation was successful then rest of error will be here.
	PostErrors error `json:"post_error"`
}

// ErrorFromSchemaError extracts the error from a SchemaError.
// It returns nil if the SchemaError pointer is nil, otherwise returns the underlying error message.
func ErrorFromSchemaError(e *SchemaError) error {
	if e == nil {
		return nil
	}
	return e.Err
}

// ErrorFromSchemaErrorList converts a SchemaErrorList into a single error by joining all errors.
// It returns nil if the SchemaErrorList is nil. All non-nil errors from the list are combined
// using errors.Join into a single error value. Returns the combined error or nil if the list is empty.
func ErrorFromSchemaErrorList(e SchemaErrors) error {
	if e.Errors == nil {
		return nil
	}

	var err error = errors.New("")

	for _, e := range e.Errors {

		if e == nil {
			continue
		}
		err = errors.Join(err, e.Err)
	}
	return err
}
