package tests_test

import (
	"bytes"
	"testing"

	"github.com/mrbns/valgo/lib/v"
)

type TestSchema struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (s *TestSchema) Rules() v.Pipeset {
	return v.NewPipesMap(v.PipeMap{
		"name": v.StringPipe(s.Name, v.NotEmpty(), v.MaxLength(10)),
		"age":  v.IntPipe(s.Age, v.NonZero()),
	})
}

func main() {

	schema := new(TestSchema)
	data := []byte(`{"name": "John", "age": 30}`)

	v.ParseBytes(data, schema)
}

func TestParse(t *testing.T) {
	data := []byte(`{"name": "John", "age": 30}`)
	schema := &TestSchema{}

	err := v.ParseBytes(data, schema)
	if err.ParseError != nil {
		t.Errorf("ParseBytes failed: %v", err.ParseError)
	}

	if schema.Name != "John" {
		t.Errorf("Expected name 'John', got '%s'", schema.Name)
	}

	if schema.Age != 30 {
		t.Errorf("Expected age 30, got %d", schema.Age)
	}
}

func TestParseFromReader(t *testing.T) {
	data := []byte(`{"name": "Jane", "age": 25}`)
	reader := bytes.NewReader(data)
	schema := &TestSchema{}

	err := v.Parse(reader, schema)
	if err.ParseError != nil {
		t.Errorf("Parse failed: %v", err.ParseError)
	}

	if schema.Name != "Jane" {
		t.Errorf("Expected name 'Jane', got '%s'", schema.Name)
	}

	if schema.Age != 25 {
		t.Errorf("Expected age 25, got %d", schema.Age)
	}
}

func TestParseInvalidJSON(t *testing.T) {
	data := []byte(`{invalid json}`)
	schema := &TestSchema{}

	err := v.ParseBytes(data, schema)
	if err.ParseError == nil {
		t.Errorf("Expected parse error for invalid JSON")
	}
}
