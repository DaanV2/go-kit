package ptr_test

import (
	"fmt"

	"github.com/daanv2/go-kit/ptr"
)

func ExampleTo() {
	value := 42
	pointer := ptr.To(value)
	fmt.Println(*pointer)
	// Output: 42
}

func ExampleTo_string() {
	str := "hello"
	strPtr := ptr.To(str)
	fmt.Println(*strPtr)
	// Output: hello
}

func ExampleFrom() {
	value := 42
	pointer := &value
	result := ptr.From(pointer)
	fmt.Println(result)
	// Output: 42
}

func ExampleFrom_nil() {
	var pointer *int
	result := ptr.From(pointer)
	fmt.Println(result)
	// Output: 0
}

func ExampleFromB() {
	value := 42
	pointer := &value
	result, ok := ptr.FromB(pointer)
	fmt.Printf("value: %d, ok: %t\n", result, ok)
	// Output: value: 42, ok: true
}

func ExampleFromB_nil() {
	var pointer *string
	result, ok := ptr.FromB(pointer)
	fmt.Printf("value: %q, ok: %t\n", result, ok)
	// Output: value: "", ok: false
}
