package xcontext_test

import (
	"context"
	"testing"

	"github.com/daanv2/go-kit/xcontext"
	"github.com/stretchr/testify/require"
)

type User struct {
	ID string
}

type Car struct {
	ID string
}

func Test_XContext_GetSet(t *testing.T) {
	userHandler := &xcontext.Handler[*User]{}
	carHandler := &xcontext.Handler[Car]{}

	t.Run("can set and get user", func(t *testing.T) {
		ctx := context.Background()
		user := User{
			ID: "user",
		}

		ctx = userHandler.Set(ctx, &user)
		result, ok := userHandler.Get(ctx)
		require.True(t, ok)
		require.NotNil(t, result)
		require.Equal(t, *result, user)
	})

	t.Run("can set and get car", func(t *testing.T) {
		ctx := context.Background()
		car := Car{
			ID: "car",
		}

		ctx = carHandler.Set(ctx, car)
		result, ok := carHandler.Get(ctx)
		require.True(t, ok)
		require.Equal(t, result, car)
	})

	t.Run("nothing set and nothing get user", func(t *testing.T) {
		ctx := context.Background()

		// ctx = userHandler.Set(ctx, &user)
		_, ok := userHandler.Get(ctx)
		require.False(t, ok)
	})

	t.Run("nothing set and nothing get car", func(t *testing.T) {
		ctx := context.Background()

		_, ok := carHandler.Get(ctx)
		require.False(t, ok)
	})
}
