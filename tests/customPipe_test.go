package tests

import (
	"errors"
	"testing"

	"github.com/mrbns/valgo/lib/v"
)

type schema struct {
	custom  string
	custom2 string
}

func (sc *schema) Rules() (v.PipeSet, error) {
	pipeMap := v.PipeMap{
		"custom": v.CustomPipe(sc.custom, func(val string) error {
			return nil
		}),
		"custom2": v.CustomPipe(sc.custom2, func(val string) error {
			if val == "" {
				return errors.New("cannot be nil")
			}
			return errors.New("custom 2 should have errors.")
		}),
	}
	return pipeMap, nil
}

func TestCustomPipe(t *testing.T) {
	data := []byte(`{"custom":"hello", "custom2": null }`)
	var buf schema
	err := v.ParseBytes(data, &buf)
	if err != nil {
		t.Error(err)
	}
}
