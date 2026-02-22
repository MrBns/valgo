package v

type customPipe[T any] struct {
	key   string
	fn    func(value T) error
	value T
}

func (p *customPipe[T]) Key() string {
	return p.key
}
func (p *customPipe[T]) Validate() error {
	return p.fn(p.value)
}
func (p *customPipe[T]) setKey(key string) {
	p.key = key
}

func CustomPipe[T any](value T, fn func(value T) error) *customPipe[T] {
	return &customPipe[T]{
		fn:    fn,
		value: value,
	}
}
