package v

// bytesPipeManager manages the validation pipeline for []byte values.
type bytesPipeManager struct {
	actions []BytesPipeAction
	value   []byte
	key     string
	error   error
}

// BytesPipeAction defines the interface for []byte validation actions.
// Each action can run validation logic on a []byte value and return an error if validation fails.
type BytesPipeAction interface {
	Run(v []byte) error
}

// BytesPipe creates a new validation pipe for []byte values.
// The pipe executes the provided actions in sequence during validation.
func BytesPipe(value []byte, actions ...BytesPipeAction) PipeFace {
	return &bytesPipeManager{
		value:   value,
		actions: actions,
		error:   nil,
	}
}

// setKey sets the validation key for this pipe.
// This key is used in error messages to identify which field failed validation.
func (pipe *bytesPipeManager) setKey(k string) {
	pipe.key = k
}

// Key returns the validation key associated with this pipe.
func (pipe *bytesPipeManager) Key() string {
	return pipe.key
}

// Validate runs all validation actions in sequence.
// Returns a FieldError if any action fails, otherwise returns nil.
func (pipe *bytesPipeManager) Validate() error {
	for _, action := range pipe.actions {
		if err := action.Run(pipe.value); err != nil {
			return NewPipeError(pipe.key, err)
		}
	}
	return nil
}
