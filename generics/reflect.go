package generics

import "reflect"

// NameOf returns the type name of the given object
func NameOf[T any]() string {
	t := reflect.TypeFor[T]()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.Name()
}

// PackageOf returns the type package of the given object
func PackageOf[T any]() string {
	return reflect.TypeFor[T]().PkgPath()
}

// SizeOf returns the amount of bytes the object is.
// NOTE: it does not traverse pointers to determine their size
func SizeOf[T any]() uintptr {
	return reflect.TypeFor[T]().Size()
}

func Empty[T any]() T {
	var empty T
	return empty
}
