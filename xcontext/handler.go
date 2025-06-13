package xcontext

import (
	"context"

	"github.com/daanv2/go-kit/generics"
)

type key[T any] struct {}

type Handler[T any] struct {
	key key[T]
}

func (h *Handler[T]) Set(ctx context.Context, v T) context.Context {
	return context.WithValue(ctx, h.key, v)
}

func (h *Handler[T]) Get(ctx context.Context) (T, bool) {
	item := ctx.Value(h.key)
	if item == nil {
		return generics.Empty[T](), false
	}

	v, ok := item.(T)
	return v, ok
}
