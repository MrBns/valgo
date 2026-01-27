package v

type floatPipeManager struct {
	actions []FloatPipeAction
	value   float64
	key     string
	error   error
}

type FloatPipeAction interface {
	Run(v float64) error
}

func NewFloatPipe(value float64, actions ...FloatPipeAction) Pipe {
	return &floatPipeManager{
		value:   value,
		actions: actions,
		error:   nil,
	}
}

func (pipe *floatPipeManager) setKey(k string) {
	pipe.key = k
}

func (pipe *floatPipeManager) Key() string {
	return pipe.key
}

func (pipe *floatPipeManager) Validate() {
	for _, action := range pipe.actions {
		if err := action.Run(pipe.value); err != nil {
			pipe.error = err
			break
		}
	}
}
