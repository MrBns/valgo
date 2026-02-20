package v

// PipeMap is a map of pipe keys to pipes.
type PipeMap map[string]PipeFace

func (m PipeMap) ValidateAll() *SchemaErrors {
	return m.validateAllSequential()
}

func (m PipeMap) validateAllSequential() *SchemaErrors {
	var errors []*SchemaError

	for key, pipe := range m {
		if err := pipe.Validate(); err != nil {
			err.Key = key
			errors = append(errors, err)
		}
	}

	if len(errors) == 0 {
		return nil
	}

	return &SchemaErrors{Errors: errors}
}

func (m PipeMap) Validate() *SchemaError {
	for key, pipe := range m {
		if err := pipe.Validate(); err != nil {
			return &SchemaError{
				Key: key,
				Err: err.Err,
			}
		}
	}
	return nil
}
