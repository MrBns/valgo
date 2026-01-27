package v

type stringPipeManager struct {
	actions []StringPipeAction
	value   string
	key     string
	error   error
}

type StringPipeAction interface {
	Run(v string) error
}

func NewStringPipe(value string, actions ...StringPipeAction) Pipe {
	return &stringPipeManager{
		value:   value,
		actions: actions,
		error:   nil,
	}
}

func (pipe *stringPipeManager) setKey(k string) {
	pipe.key = k
}

func (pipe *stringPipeManager) Key() string {
	return pipe.key
}

func (pipe *stringPipeManager) Validate() {

	// hasError := false

	for _, action := range pipe.actions {
		if err := action.Run(pipe.value); err != nil {
			// hasError = true
			pipe.error = err
			break
		}
	}

}
