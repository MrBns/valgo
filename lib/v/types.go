package v

type Schema interface {
	Validate() []SchemaValidationError
}

type SchemaValidationError struct {
	Key string `json:"key"`
	Msg string `json:"msg"`
}

type Pipe interface {
	Key() string
	Validate()
	setKey(string)
}

type PipeAction interface {
	Run(v any) error
}

type PipeActionOption interface {
	Run(v any) error
}

type ActionOptions interface {
	Run(v any) error
}

type ErrMsg interface {
	ActionOptions
	Msg(v ...any) string
}
