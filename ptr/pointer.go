package ptr

import "github.com/daanv2/go-kit/generics"

func To[T any](v T) *T {
	return &v
}

func From[T any](v *T) T {
	if v == nil {
		return generics.Empty[T]()
	}

	return *v
}

func FromB[T any](v *T) (T, bool) {
	if v == nil {
		return generics.Empty[T](), false
	}

	return *v, true
}