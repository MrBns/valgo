// Package v provides a flexible validation framework for Go.
// It supports validation of strings, integers, and floats through a pipeline pattern.
//
// Example usage:
//
//	pipe := v.NewStringPipe("test@example.com", v.IsEmail())
//	if err := pipe.Validate(); err != nil {
//	    log.Fatal(err)
//	}
package v
