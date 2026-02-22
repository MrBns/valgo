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
