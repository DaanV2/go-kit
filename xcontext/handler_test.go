package xcontext_test

import (
	"context"
	"fmt"
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

// Examples

func ExampleHandler_Set() {
	type User struct {
		Name string
	}

	handler := &xcontext.Handler[User]{}
	ctx := context.Background()

	user := User{Name: "Alice"}
	ctx = handler.Set(ctx, user)

	retrieved, ok := handler.Get(ctx)
	if ok {
		fmt.Println(retrieved.Name)
	}
	// Output: Alice
}

func ExampleHandler_Get() {
	type RequestID string

	handler := &xcontext.Handler[RequestID]{}
	ctx := context.Background()

	// Get from empty context
	_, ok := handler.Get(ctx)
	fmt.Printf("Value found: %t\n", ok)

	// Set and get
	ctx = handler.Set(ctx, RequestID("req-123"))
	value, ok := handler.Get(ctx)
	fmt.Printf("Value found: %t, Value: %s\n", ok, value)
	// Output: Value found: false
	// Value found: true, Value: req-123
}

func ExampleHandler_Set_multipleTypes() {
	type UserID int
	type SessionID string

	userHandler := &xcontext.Handler[UserID]{}
	sessionHandler := &xcontext.Handler[SessionID]{}

	ctx := context.Background()
	ctx = userHandler.Set(ctx, UserID(42))
	ctx = sessionHandler.Set(ctx, SessionID("sess-abc"))

	userID, _ := userHandler.Get(ctx)
	sessionID, _ := sessionHandler.Get(ctx)

	fmt.Printf("User: %d, Session: %s\n", userID, sessionID)
	// Output: User: 42, Session: sess-abc
}
