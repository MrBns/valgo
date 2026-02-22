package tests_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/mrbns/valgo/lib/v"
)

func TestSchemaErrorMessageIsMeaningful(t *testing.T) {
	err := v.NewPipeError("name", errors.New("cannot be empty"))
	if err == nil {
		t.Fatalf("expected schema error")
	}
	if got := err.Error(); got != "name: cannot be empty" {
		t.Fatalf("unexpected error message: %q", got)
	}
}

func TestSchemaErrorsJoinMessages(t *testing.T) {
	errList := v.ValidationErrors{
		v.NewPipeError("name", errors.New("cannot be empty")),
		v.NewPipeError("age", errors.New("value must be non-zero")),
	}

	if len(errList) == 0 {
		t.Fatalf("expected schema errors")
	}

	msg := errList.Error()
	if !strings.Contains(msg, "cannot be empty") || !strings.Contains(msg, "value must be non-zero") {
		t.Fatalf("joined message should include both errors, got: %q", msg)
	}
}

func TestParseBytesReturnsNilOnSuccess(t *testing.T) {
	data := []byte(`{"name":"John","age":30}`)
	schema := &TestSchema{}

	err := v.ParseBytes(data, schema)
	if err != nil {
		t.Fatalf("expected nil parse error, got: %v", err)
	}
}

func TestParseBytesFullReturnsNilOnSuccess(t *testing.T) {
	data := []byte(`{"name":"John","age":30}`)
	schema := &TestSchema{}

	err := v.ParseBytesFull(data, schema)
	if err != nil {
		t.Fatalf("expected nil parse error, got: %v", err)
	}
}

func TestParseErrorsMessagePrefix(t *testing.T) {
	data := []byte(`{invalid json}`)
	schema := &TestSchema{}

	err := v.ParseBytesFull(data, schema)
	if err == nil {
		t.Fatalf("expected parse error")
	}
	if !strings.HasPrefix(err.Error(), "v.parse_error: ") {
		t.Fatalf("unexpected parse error format: %q", err.Error())
	}
}

func TestErrorIsAndAs(t *testing.T) {
	// 1. Test errors.As with ParseError
	data := []byte(`{invalid json}`)
	schema := &TestSchema{}

	err := v.ParseBytesFull(data, schema)
	if err == nil {
		t.Fatalf("expected parse error")
	}

	var parseErr *v.ParseError
	if !errors.As(err, &parseErr) {
		t.Fatalf("expected error to be of type *v.ParseError")
	}
	if parseErr.ParseError == nil {
		t.Fatalf("expected ParseError field to be populated")
	}

	// 2. Test errors.As with ValidationErrors
	data = []byte(`{"name": "", "age": 0}`) // Invalid data based on TestSchema rules
	err = v.ParseBytesFull(data, schema)
	if err == nil {
		t.Fatalf("expected validation error")
	}

	if !errors.As(err, &parseErr) {
		t.Fatalf("expected error to be of type *v.ParseError")
	}

	var validationErrs v.ValidationErrors
	if !errors.As(parseErr.ValidationError, &validationErrs) {
		t.Fatalf("expected ValidationError to be of type v.ValidationErrors")
	}

	if len(validationErrs) != 2 {
		t.Fatalf("expected 2 validation errors, got %d", len(validationErrs))
	}

	// 3. Test errors.Is with specific underlying errors
	// Let's create a specific error to test errors.Is
	specificErr := errors.New("a very specific error")
	pipeErr := v.NewPipeError("custom_field", specificErr)

	if !errors.Is(pipeErr, specificErr) {
		t.Fatalf("expected errors.Is to find the specific underlying error")
	}

	// Test errors.Is with ValidationErrors slice
	errList := v.ValidationErrors{
		v.NewPipeError("field1", errors.New("error 1")),
		v.NewPipeError("field2", specificErr),
	}

	if !errors.Is(errList, specificErr) {
		t.Fatalf("expected errors.Is to find the specific error inside the ValidationErrors slice")
	}
}
