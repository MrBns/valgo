package v

// floatPipeManager manages the validation pipeline for float64 values.
type floatPipeManager struct {
	actions []FloatPipeAction
	value   float64
	key     string
	error   error
}

// FloatPipeAction defines the interface for float64 validation actions.
// Each action can run validation logic on a float64 value and return an error if validation fails.
type FloatPipeAction interface {
	Run(v float64) error
}

// NewFloatPipe creates a new validation pipe for float64 values.
// The pipe executes the provided actions in sequence during validation.
//
// Example:
//
//	pipe := NewFloatPipe(42.5, MinFloat(0), MaxFloat(100))
func NewFloatPipe(value float64, actions ...FloatPipeAction) Pipe {
	return &floatPipeManager{
		value:   value,
		actions: actions,
		error:   nil,
	}
}

// setKey sets the validation key for this pipe.
// This key is used in error messages to identify which field failed validation.
func (pipe *floatPipeManager) setKey(k string) {
	pipe.key = k
}

// Key returns the validation key associated with this pipe.
func (pipe *floatPipeManager) Key() string {
	return pipe.key
}

// Validate runs all validation actions in sequence.
// Returns a SchemaValidationError if any action fails, otherwise returns nil.
func (pipe *floatPipeManager) Validate() *SchemaValidationError {
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
