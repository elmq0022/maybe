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

func Map[T, U any](m Maybe[T], f func(T) U) Maybe[U] {
	if m.valid {
		return Some(f(m.value))
	}
	return None[U]()
}

func (m *Maybe[T]) Unwrap() T {
	if !m.valid {
		panic("called Unwrap on None val")
	}
	return m.value
}

func (m *Maybe[T]) UnwrapOr(def T) T {
	if m.valid {
		return m.value
	}
	return def
}
