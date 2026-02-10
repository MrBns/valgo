package v

type Schema interface {
	Rules() (Pipeset, error)
}

// Pipeset is the interface that wraps the Validate method.
// Validate returns a list of SchemaValidationError if validation fails.
type Pipeset interface {
	ValidateAll() *SchemaErrors
	Validate() *SchemaError
}

// PipeFace is the interface for a validation pipe.
// a pipe is a sequence of actions that validates a value.
type PipeFace interface {
	Key() string
	Validate() *SchemaError
	setKey(string)
}

// PipeActionFace is the interface for a pipe action.
// an action is a single validation step.
type PipeActionFace interface {
	Run(v any) error
}

// PipeActionOptionFace is the interface for a pipe action option.
// an option is a configuration for a pipe action.
type PipeActionOptionFace interface {
	Run(v any) error
}

// ActionOptionFace is the interface for options passed to actions.
type ActionOptionFace interface {
	Run(v any) error
}

// CustomErrFace is the interface for custom error messages.
// note that [CustomErrFace] is also a [ActionOptionFace].
type CustomErrFace interface {
	ActionOptionFace
	Msg(v any) string
}
