
# maybe

A small, zero-dependency generic Maybe/Option container for Go (1.18+).

This package provides a simple `Maybe[T]` type that explicitly represents a value that may be present (`Some`) or absent (`None`). It aims to make code that deals with optional values clearer and less error-prone than using raw pointers or sentinel values.

Highlights

- Small and focused: a handful of helpers (`Some`, `None`, `Map`, `Unwrap`, `UnwrapOr`, and checks).
- Uses Go generics; works with any type `T`.
- No runtime allocations beyond what your value already requires.

Why use this?

Pointers are a common way to encode optional values in Go, but they can be noisy and less self-documenting. `Maybe[T]` makes intent explicit and provides small, composable helpers for transforming and working with optional values.

Example

```go
package main

import (
    "fmt"

    m "github.com/elmq0022/maybe/maybe"
)

func main() {
    a := m.Some(42)
    b := m.None[int]()

    if a.IsSome() {
        fmt.Println("a is some:", a.Unwrap())
    }

    // map over a Maybe
    s := m.Map(a, func(v int) string { return fmt.Sprintf("%d apples", v) })
    fmt.Println(s.Unwrap()) // "42 apples"

    // fallback/default
    fmt.Println(b.UnwrapOr(100)) // prints 100
}
```

API (short)

- `Maybe[T]` — container type (fields are unexported).
- `Some[T](value T) Maybe[T]` — construct a present value.
- `None[T]() Maybe[T]` — construct an empty value.
- `(m *Maybe[T]) IsSome() bool` — true when value is present.
- `(m *Maybe[T]) IsNone() bool` — true when value is absent.
- `Map[T,U](m Maybe[T], f func(T) U) Maybe[U]` — map a function over a maybe.
- `(m *Maybe[T]) Unwrap() T` — returns the inner value or panics if `None`.
- `(m *Maybe[T]) UnwrapOr(def T) T` — returns the inner value or the provided default.

Installation

This package uses Go modules. From your project root:

```bash
go get github.com/elmq0022/maybe/maybe
```

Running tests

From the project root:

```bash
go test ./...
```

Contributing

This repository is provided for public use (you may import and use the code freely), but it is currently closed to external contributions. Please do not open pull requests — issues and PRs will not be accepted at this time.

If you have questions about using the package, feel free to open an issue to ask for clarification; maintainers may respond on a best-effort basis.

License

This project is provided under the terms in `LICENSE`.
