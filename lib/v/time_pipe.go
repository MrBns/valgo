package v

import "time"

// timePipeManager manages the validation pipeline for time.Time values.
type timePipeManager struct {
	actions []TimePipeAction
	value   time.Time
	key     string
	error   error
}

// TimePipeAction defines the interface for time validation actions.
// Each action can run validation logic on a time.Time value and return an error if validation fails.
type TimePipeAction interface {
	Run(v time.Time) error
}

// TimePipe creates a new validation pipe for time.Time values.
// The pipe executes the provided actions in sequence during validation.
//
// Example:
//
//	pipe := TimePipe(time.Now(), BeforeNow(), After(time.Now().AddDate(0, 0, -7)))
func TimePipe(value time.Time, actions ...TimePipeAction) PipeFace {
	return &timePipeManager{
		value:   value,
		actions: actions,
		error:   nil,
	}
}

// setKey sets the validation key for this pipe.
// This key is used in error messages to identify which field failed validation.
func (pipe *timePipeManager) setKey(k string) {
	pipe.key = k
}

// Key returns the validation key associated with this pipe.
func (pipe *timePipeManager) Key() string {
	return pipe.key
}

// Validate runs all validation actions in sequence.
// Returns a SchemaValidationError if any action fails, otherwise returns nil.
func (pipe *timePipeManager) Validate() *SchemaError {
	for _, action := range pipe.actions {
		if err := action.Run(pipe.value); err != nil {
			return &SchemaError{
				Key: pipe.key,
				Err: err,
			}
		}
	}
	return nil
}
