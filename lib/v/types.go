package v

// Schema is the interface that wraps the Validate method.
// Validate returns a list of SchemaValidationError if validation fails.
type Schema interface {
	Validate() []SchemaValidationError
}

// SchemaValidationError represents a validation error with a key and an error message.
type SchemaValidationError struct {
	Key string `json:"key"`
	Err error  `json:"msg"`
}

// Pipe is the interface for a validation pipe.
// a pipe is a sequence of actions that validates a value.
type Pipe interface {
	Key() string
	Validate() *SchemaValidationError
	setKey(string)
}

// PipeAction is the interface for a pipe action.
// an action is a single validation step.
type PipeAction interface {
	Run(v any) error
}

// PipeActionOption is the interface for a pipe action option.
// an option is a configuration for a pipe action.
type PipeActionOption interface {
	Run(v any) error
}

// ActionOptions is the interface for options passed to actions.
type ActionOptions interface {
	Run(v any) error
}

// ErrMsg is the interface for custom error messages.
type ErrMsg interface {
	ActionOptions
	Msg(v ...any) string
}
