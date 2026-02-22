package v

// PipeMap is a map of pipe keys to pipes.
type PipeMap map[string]PipeFace

func (m PipeMap) ValidateAll() error {
	return m.validateAllSequential()
}

func (m PipeMap) validateAllSequential() error {
	var validationErrors ValidationErrors

	for key, pipe := range m {
		if err := pipe.Validate(); err != nil {
			if fieldErr, ok := err.(*PipeError); ok {
				validationErrors = append(validationErrors, fieldErr)
			} else {
				validationErrors = append(validationErrors, NewPipeError(key, err))
			}
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}
	return nil
}

func (m PipeMap) Validate() error {
	for key, pipe := range m {
		if err := pipe.Validate(); err != nil {
			if fieldErr, ok := err.(*PipeError); ok {
				return fieldErr
			}
			return NewPipeError(key, err)
		}
	}
	return nil
}
