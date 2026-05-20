package tests_test

import (
	"testing"

	"github.com/mrbns/valgo/lib/v"
)

func TestBytesNotEmpty(t *testing.T) {
	err := v.BytesNotEmpty().Run([]byte("hello"))
	if err != nil {
		t.Error(err)
	}

	err = v.BytesNotEmpty().Run([]byte{})
	if err == nil {
		t.Errorf("empty bytes should throw error")
	}
}

func TestBytesLength(t *testing.T) {
	err := v.MinBytesLength(3).Run([]byte("hello"))
	if err != nil {
		t.Error(err)
	}

	err = v.MinBytesLength(10).Run([]byte("hello"))
	if err == nil {
		t.Errorf("short bytes should throw error")
	}

	err = v.MaxBytesLength(10).Run([]byte("hello"))
	if err != nil {
		t.Error(err)
	}

	err = v.MaxBytesLength(3).Run([]byte("hello"))
	if err == nil {
		t.Errorf("long bytes should throw error")
	}
}

func TestEqualBytes(t *testing.T) {
	err := v.EqualBytes([]byte("hello")).Run([]byte("hello"))
	if err != nil {
		t.Error(err)
	}

	err = v.EqualBytes([]byte("hello")).Run([]byte("world"))
	if err == nil {
		t.Errorf("different bytes should throw error")
	}
}

func TestContainsBytes(t *testing.T) {
	err := v.ContainsBytes([]byte("ell")).Run([]byte("hello"))
	if err != nil {
		t.Error(err)
	}

	err = v.ContainsBytes([]byte("zzz")).Run([]byte("hello"))
	if err == nil {
		t.Errorf("missing bytes should throw error")
	}
}

func TestBytesPipe(t *testing.T) {
	pipe := v.BytesPipe([]byte("payload"), v.BytesNotEmpty(), v.MinBytesLength(3))
	if err := pipe.Validate(); err != nil {
		t.Errorf("BytesPipe should pass for valid bytes: %v", err)
	}

	pipe = v.BytesPipe([]byte{}, v.BytesNotEmpty())
	if err := pipe.Validate(); err == nil {
		t.Errorf("BytesPipe should fail for empty bytes")
	}
}
