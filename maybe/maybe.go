package maybe

type Maybe[T any] struct {
	value T
	valid bool
}

func Some[T any](value T) Maybe[T] {
	return Maybe[T]{value: value, valid: true}
}

func None[T any]() Maybe[T] {
	return Maybe[T]{valid: false}
}

func (m *Maybe[T]) IsSome() bool {
	return m.valid
}

func (m *Maybe[T]) IsNone() bool {
	return !m.valid
}
