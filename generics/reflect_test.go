package generics_test

import (
	"fmt"
	"testing"

	"github.com/daanv2/go-kit/generics"
	"github.com/stretchr/testify/require"
)

type GenericTypeTest struct {
	id int32
}

func Test_NameOf(t *testing.T) {
	name := generics.NameOf[GenericTypeTest]()

	require.Equal(t, name, "GenericTypeTest")
}

func Test_PackageOf(t *testing.T) {
	pkg := generics.PackageOf[GenericTypeTest]()

	require.Equal(t, pkg, "github.com/daanv2/go-kit/generics_test")
}

func Test_SizeOf(t *testing.T) {
	size := generics.SizeOf[GenericTypeTest]()

	require.EqualValues(t, size, 4)
}

func Test_Empty(t *testing.T) {
	empty := generics.Empty[GenericTypeTest]()

	require.EqualValues(t, empty.id, 0)
}

func Test_NameOf_Table(t *testing.T) {
	type nameTestCase struct {
		desc     string
		expected string
		fn       func() string
	}
	tests := []nameTestCase{
		{"bool", "bool", generics.NameOf[bool]},
		{"uint", "uint", generics.NameOf[uint]},
		{"uint8", "uint8", generics.NameOf[uint8]},
		{"uint16", "uint16", generics.NameOf[uint16]},
		{"uint32", "uint32", generics.NameOf[uint32]},
		{"uint64", "uint64", generics.NameOf[uint64]},
		{"int", "int", generics.NameOf[int]},
		{"int8", "int8", generics.NameOf[int8]},
		{"int16", "int16", generics.NameOf[int16]},
		{"int32", "int32", generics.NameOf[int32]},
		{"int64", "int64", generics.NameOf[int64]},
		{"float32", "float32", generics.NameOf[float32]},
		{"float64", "float64", generics.NameOf[float64]},
		{"complex128", "complex128", generics.NameOf[complex128]},
		{"byte (alias for uint8)", "uint8", generics.NameOf[byte]},
		{"rune (alias for int32)", "int32", generics.NameOf[rune]},
		{"any", "", generics.NameOf[any]},
		{"TestCase", "TestCase", generics.NameOf[TestCase]},

		// Pointers
		{"*bool", "bool", generics.NameOf[*bool]},
		{"*uint", "uint", generics.NameOf[*uint]},
		{"*uint8", "uint8", generics.NameOf[*uint8]},
		{"*uint16", "uint16", generics.NameOf[*uint16]},
		{"*uint32", "uint32", generics.NameOf[*uint32]},
		{"*uint64", "uint64", generics.NameOf[*uint64]},
		{"*int", "int", generics.NameOf[*int]},
		{"*int8", "int8", generics.NameOf[*int8]},
		{"*int16", "int16", generics.NameOf[*int16]},
		{"*int32", "int32", generics.NameOf[*int32]},
		{"*int64", "int64", generics.NameOf[*int64]},
		{"*float32", "float32", generics.NameOf[*float32]},
		{"*float64", "float64", generics.NameOf[*float64]},
		{"*complex128", "complex128", generics.NameOf[*complex128]},
		{"*byte (alias for uint8)", "uint8", generics.NameOf[*byte]},
		{"*rune (alias for int32)", "int32", generics.NameOf[*rune]},
		{"*TestCase", "TestCase", generics.NameOf[*TestCase]},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.fn())
		})
	}
}

func Test_SizeOf_Table(t *testing.T) {
	type sizeTestCase struct {
		desc     string
		expected uintptr
		fn       func() uintptr
	}
	tests := []sizeTestCase{
		{"bool", 1, generics.SizeOf[bool]},
		{"uint8", 1, generics.SizeOf[uint8]},
		{"uint16", 2, generics.SizeOf[uint16]},
		{"uint32", 4, generics.SizeOf[uint32]},
		{"uint64", 8, generics.SizeOf[uint64]},
		{"int8", 1, generics.SizeOf[int8]},
		{"int16", 2, generics.SizeOf[int16]},
		{"int32", 4, generics.SizeOf[int32]},
		{"int64", 8, generics.SizeOf[int64]},
		{"float32", 4, generics.SizeOf[float32]},
		{"float64", 8, generics.SizeOf[float64]},
		{"complex128", 16, generics.SizeOf[complex128]},
		{"byte (alias for uint8)", 1, generics.SizeOf[byte]},
		{"rune (alias for int32)", 4, generics.SizeOf[rune]},
		{"any", 16, generics.SizeOf[any]},
		{"TestCase", 12, generics.SizeOf[TestCase]},
		{"*TestCase", 8, generics.SizeOf[*TestCase]},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			require.EqualValues(t, tc.expected, tc.fn())
		})
	}
}

type TestCase struct {
	A int32
	B int32
	C int32
}

// Examples

func ExampleNameOf() {
	name := generics.NameOf[int]()
	fmt.Println(name)
	// Output: int
}

func ExampleNameOf_customType() {
	name := generics.NameOf[TestCase]()
	fmt.Println(name)
	// Output: TestCase
}

func ExampleNameOf_pointer() {
	name := generics.NameOf[*string]()
	fmt.Println(name)
	// Output: string
}

func ExamplePackageOf() {
	pkg := generics.PackageOf[TestCase]()
	fmt.Println(pkg)
	// Output: github.com/daanv2/go-kit/generics_test
}

func ExamplePackageOf_builtin() {
	pkg := generics.PackageOf[int]()
	fmt.Println(pkg)
	// Output: 
}

func ExampleSizeOf() {
	size := generics.SizeOf[int32]()
	fmt.Println(size)
	// Output: 4
}

func ExampleSizeOf_struct() {
	size := generics.SizeOf[TestCase]()
	fmt.Println(size)
	// Output: 12
}

func ExampleEmpty() {
	empty := generics.Empty[int]()
	fmt.Println(empty)
	// Output: 0
}

func ExampleEmpty_struct() {
	type Person struct {
		Name string
		Age  int
	}
	empty := generics.Empty[Person]()
	fmt.Printf("Name: %q, Age: %d\n", empty.Name, empty.Age)
	// Output: Name: "", Age: 0
}
