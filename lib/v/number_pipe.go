package v

type numberPipeManager struct {
	actions []NumberPipeAction
	value   int64
	key     string
	error   error
}

type NumberPipeAction interface {
	Run(v int64) error
}

func NewNumberPipe(value int64, actions ...NumberPipeAction) Pipe {
	return &numberPipeManager{
		value:   value,
		actions: actions,
		error:   nil,
	}
}

func (pipe *numberPipeManager) setKey(k string) {
	pipe.key = k
}

func (pipe *numberPipeManager) Key() string {
	return pipe.key
}

func (pipe *numberPipeManager) Validate() {
	for _, action := range pipe.actions {
		if err := action.Run(pipe.value); err != nil {
			pipe.error = err
			break
		}
	}
}
