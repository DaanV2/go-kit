package generics

import "reflect"

// NameOf returns the type name of the given object
func NameOf[T any]() string {
	var item T
	return reflect.TypeOf(item).Name()
}

// PackageOf returns the type package of the given object
func PackageOf[T any]() string {
	var item T
	return reflect.TypeOf(item).PkgPath()
}

// SizeOf returns the amount of bytes the object is.
// NOTE: it does not traverse pointers to determine their size
func SizeOf[T any]() uint {
	var item T
	return uint(reflect.TypeOf(item).Size())
}

func Empty[T any]() T {
	var empty T
	return empty
}