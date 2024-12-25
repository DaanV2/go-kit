package generics_test

import (
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
