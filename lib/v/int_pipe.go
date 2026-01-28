package v

// IntPipeManager manages the validation pipeline for int64 values.
type IntPipeManager struct {
	actions []IntPipeAction
	value   int64
	key     string
	error   error
}

// IntPipeAction defines the interface for int64 validation actions.
// Each action can run validation logic on an int64 value and return an error if validation fails.
type IntPipeAction interface {
	Run(v int64) error
}

// NewNumberPipe creates a new validation pipe for int64 values.
// The pipe executes the provided actions in sequence during validation.
//
// Example:
//
//	pipe := NewNumberPipe(42, Min(0), Max(100))
func NewNumberPipe(value int64, actions ...IntPipeAction) Pipe {
	return &IntPipeManager{
		value:   value,
		actions: actions,
		error:   nil,
	}
}

// setKey sets the validation key for this pipe.
// This key is used in error messages to identify which field failed validation.
func (pipe *IntPipeManager) setKey(k string) {
	pipe.key = k
}

// Key returns the validation key associated with this pipe.
func (pipe *IntPipeManager) Key() string {
	return pipe.key
}

// Validate runs all validation actions in sequence.
// Returns a SchemaValidationError if any action fails, otherwise returns nil.
func (pipe *IntPipeManager) Validate() *SchemaValidationError {
	for _, action := range pipe.actions {
		if err := action.Run(pipe.value); err != nil {
			return &SchemaValidationError{
				Key: pipe.key,
				Err: err,
			}
		}
	}
	return nil
}
