package v

// stringPipeManager manages the validation pipeline for string values.
type stringPipeManager struct {
	actions []StringPipeAction
	value   string
	key     string
	error   error
}

// StringPipeAction defines the interface for string validation actions.
// Each action can run validation logic on a string value and return an error if validation fails.
type StringPipeAction interface {
	Run(v string) error
}

// StringPipe creates a new validation pipe for string values.
// The pipe executes the provided actions in sequence during validation.
//
// Example:
//
//	pipe := StringPipe("user@example.com", Empty(), IsEmail())
//	if err := pipe.Validate(); err != nil {
//	    log.Fatal(err)
//	}
func StringPipe(value string, actions ...StringPipeAction) PipeFace {
	return &stringPipeManager{
		value:   value,
		actions: actions,
		error:   nil,
	}
}

// setKey sets the validation key for this pipe.
// This key is used in error messages to identify which field failed validation.
func (pipe *stringPipeManager) setKey(k string) {
	pipe.key = k
}

// Key returns the validation key associated with this pipe.
func (pipe *stringPipeManager) Key() string {
	return pipe.key
}

// Validate runs all validation actions in sequence.
// Returns a SchemaValidationError if any action fails, otherwise returns nil.
func (pipe *stringPipeManager) Validate() *SchemaError {

	// hasError := false

	for _, action := range pipe.actions {
		if err := action.Run(pipe.value); err != nil {
			// hasError = true
			return &SchemaError{
				Key: pipe.key,
				Err: err,
			}
		}

	}

	return nil
}
