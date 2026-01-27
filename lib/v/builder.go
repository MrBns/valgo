package v

type PipeRegistry struct {
	pipes []Pipe
}

func NewPipesBuilder(pipeEntries ...Pipe) Schema {
	return &PipeRegistry{
		pipes: pipeEntries,
	}
}

func (schema *PipeRegistry) Validate() []SchemaValidationError {

	for _, pipe := range schema.pipes {
		pipe.Validate()
	}

	return nil
}

type PipeMap map[string]Pipe

func NewPipesMap(pipeMap PipeMap) Schema {
	pipes := []Pipe{}
	for k, v := range pipeMap {
		v.setKey(k)
		pipes = append(pipes, v)
	}
	return &PipeRegistry{
		pipes: pipes,
	}
}

/* Key Builder */
type pipeEntry struct {
	key string
}

func Entry(key string) *pipeEntry {
	return &pipeEntry{
		key: key,
	}
}

func (pk *pipeEntry) StringPipe(value string, actions ...StringPipeAction) Pipe {
	return &stringPipeManager{
		value:   value,
		actions: actions,
		error:   nil,
	}
}
