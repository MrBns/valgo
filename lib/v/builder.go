package v

import (
	"strconv"
	"strings"
	"sync"
)

// PipeRegistry where all pipes will be compiled and saved.
// it will be the final object
type PipeRegistry struct {
	pipes []PipeFace
}

// NewPipesBuilder creates a new [SchemaFace] with the given pipes.
func NewPipesBuilder(pipeEntries ...PipeFace) SchemaFace {
	return &PipeRegistry{
		pipes: pipeEntries,
	}
}

// ValidateAll returns array of [SchemaError] but if there is no error then it returns nil.
//
// it validate all the pipes. but return the first error that pipe. but pipe will be ignored if there is no error.
func (schema *PipeRegistry) ValidateAll() SchemaErrorList {

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	errs := SchemaErrorList{}

	for _, pipe := range schema.pipes {
		wg.Add(1)
		go func(p PipeFace) {
			defer wg.Done()

			if err := p.Validate(); err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
			}
		}(pipe)
	}

	wg.Wait()
	if len(errs) == 0 {
		return nil
	}
	return errs
}

func (schema *PipeRegistry) Validate() *SchemaError {
	for _, pipe := range schema.pipes {
		if err := pipe.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// PipeMap is a map of pipe keys to pipes.
type PipeMap map[string]PipeFace

// NewPipesMap creates a new [SchemaFace] from a [PipeMap].
func NewPipesMap(pipeMap PipeMap) SchemaFace {
	pipes := []PipeFace{}
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
func (pk *PipeEntryKeyHolder) StringPipe(value string, actions ...StringPipeAction) PipeFace {
	return &stringPipeManager{
		value:   value,
		actions: actions,
		error:   nil,
	}
}

// IntPipe Creates a Int Pipe
func (pk *PipeEntryKeyHolder) IntPipe(value int, actions ...IntPipeAction) PipeFace {
	return &IntPipeManager{
		value:   value,
		actions: actions,
		error:   nil,
	}
}

// FloatPipe Creates a Float Pipe
func (pk *PipeEntryKeyHolder) FloatPipe(value float64, actions ...FloatPipeAction) PipeFace {
	return &floatPipeManager{
		value:   value,
		actions: actions,
		error:   nil,
	}
}

// Custom Error message
type CustomErrMsg struct {
	msg string
}

// A getter to Get Error message.
func (c *CustomErrMsg) Msg(v any) string {

	msg := c.msg

	if !strings.Contains(msg, "{VALUE}") {
		return msg
	}

	value := ""

	switch val := v.(type) {
	case string:
		value = val
	case int:
		value = strconv.Itoa(val)
	case int8:
		value = strconv.FormatInt(int64(val), 10)
	case int16:
		value = strconv.FormatInt(int64(val), 10)
	case int32:
		value = strconv.FormatInt(int64(val), 10)
	case int64:
		value = strconv.FormatInt(val, 10)
	case uint:
		value = strconv.FormatUint(uint64(val), 10)
	case uint8:
		value = strconv.FormatUint(uint64(val), 10)
	case uint16:
		value = strconv.FormatUint(uint64(val), 10)
	case uint32:
		value = strconv.FormatUint(uint64(val), 10)
	case uint64:
		value = strconv.FormatUint(val, 10)
	case float32:
		value = strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		value = strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		value = strconv.FormatBool(val)
	default:
		return msg
	}

	return strings.Replace(msg, "{VALUE}", value, -1)
}

func (c *CustomErrMsg) Run(v any) error {
	return nil
}

// ErrMsg is used to specify custom error message
func ErrMsg(v string) CustomErrFace {
	return &CustomErrMsg{
		msg: v,
	}
}
