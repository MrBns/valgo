package v

// PipeRegistry where all pipes will be compiled and saved.
// it will be the final object
type PipeRegistry struct {
	pipes []Pipe
}

// NewPipesBuilder creates a new [Schema] with the given pipes.
func NewPipesBuilder(pipeEntries ...Pipe) Schema {
	return &PipeRegistry{
		pipes: pipeEntries,
	}
}

// Validate returns array of [SchemaValidationError] but if there is no error then it returns nil.
//
// it validate all the pipes. but return the first error that pipe. but pipe will be ignored if there is no error.
func (schema *PipeRegistry) Validate() []SchemaValidationError {

	for _, pipe := range schema.pipes {
		pipe.Validate()
	}

	return nil
}

// PipeMap is a map of pipe keys to pipes.
type PipeMap map[string]Pipe

// NewPipesMap creates a new [Schema] from a [PipeMap].
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

// PipeEntryKeyHolder just hold the key to build pipe in registry.
type PipeEntryKeyHolder struct {
	key string
}

// Entry is a constructor method that returns an [PipeEntryKeyHolder] to build pipe
func Entry(key string) *PipeEntryKeyHolder {
	return &PipeEntryKeyHolder{
		key: key,
	}
}

// StringPipe Creates a String Pipe
func (pk *PipeEntryKeyHolder) StringPipe(value string, actions ...StringPipeAction) Pipe {
	return &stringPipeManager{
		value:   value,
		actions: actions,
		error:   nil,
	}
}

// IntPipe Creates a Int Pipe
func (pk *PipeEntryKeyHolder) IntPipe(value int64, actions ...IntPipeAction) Pipe {
	return &IntPipeManager{
		value:   value,
		actions: actions,
		error:   nil,
	}
}

// FloatPipe Creates a Float Pipe
func (pk *PipeEntryKeyHolder) FloatPipe(value float64, actions ...FloatPipeAction) Pipe {
	return &floatPipeManager{
		value:   value,
		actions: actions,
		error:   nil,
	}
}
