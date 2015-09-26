// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package generic

import (
	"errors"
	"log"
	"reflect"
)

var (
	ErrNotAStruct   = errors.New("Argument does not reference a struct")
	ErrUnknownField = errors.New("Struct has no field of the given name")
)

// VerifyInt requires i to be of type int or panics.
func VerifyInt(i interface{}) {
	if !IsInt(i) {
		log.Panic("Int parameter required, not %#v (%T)\n", i, i)
	}
}

// VerifyString requires i to be of type string or panics.
func VerifyString(i interface{}) {
	if !IsString(i) {
		log.Panic("Int parameter required, not %#v (%T)\n", i, i)
	}
}

// VerifySlice requires i to be of type slice or panics.
func VerifySlice(i interface{}) {
	if IsSlice(i) {
		log.Panic("Int parameter required, not %#v (%T)\n", i, i)
	}
}

// IsInt returns true if i is of type int.
func IsInt(i interface{}) bool {
	return Indirect(i).Type().Kind() == reflect.Int
}

// IsSlice returns true if i is of type slice.
func IsSlice(i interface{}) bool {
	return Indirect(i).Type().Kind() == reflect.Slice
}

// IsPtr returns true if i is of type pointer.
func IsPtr(i interface{}) bool {
	return reflect.TypeOf(i).Kind() == reflect.Ptr
}

// IsString returns true if i is of type string.
func IsString(i interface{}) bool {
	return reflect.TypeOf(i).Kind() == reflect.String
}

// IsStruct returns true if i is of type struct.
func IsStruct(i interface{}) bool {
	return TypeIsStruct(reflect.TypeOf(i))
}

// TypeIsStruct returns true if type t is of type struct.
func TypeIsStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

// IsStruct returns true if i is of type pointer.
func IsPointer(i interface{}) bool {
	return TypeIsPointer(reflect.TypeOf(i))
}

// TypeIsPointer returns true if type t is of type pointer.
func TypeIsPointer(typ reflect.Type) bool {
	return typ != nil && typ.Kind() == reflect.Ptr
}

// IsStructPtr returns true if i is of type struct pointer.
func IsStructPtr(i interface{}) bool {
	if !IsPtr(i) {
		return false
	}
	return IsStruct(Indirect(i))
}

// IsStructOrStructPtr returns true if i is of type struct or struct pointer.
func IsStructOrStructPtr(i interface{}) bool {
	return IsStruct(Indirect(i))
}

// Indirect returns the value that i points to or,
// if i is not a pointer, returns i.
func Indirect(i interface{}) reflect.Value {
	return reflect.Indirect(reflect.ValueOf(i))
}

// Append the element(s) represented by k to the slice represented by ret
// and return the result.  Both parameters must have the same underlying type.
// If k is a slice, the elements of k are appended to ret.
func Append(ret, k interface{}) reflect.Value {
	if IsSlice(k) {
		return reflect.AppendSlice(ValueOf(ret), ValueOf(k))
	}
	return reflect.Append(ValueOf(ret), ValueOf(k))
}

// MakeSlice creates a new, empty slice value for the given type. The
// parameter is expected to be an exemplar of the slice type to make.
func MakeSlice(i interface{}) reflect.Value {
	ut := Indirect(i).Type()
	if ut.Kind() == reflect.Slice {
		return reflect.MakeSlice(ut, 0, 0)
	}
	st := reflect.SliceOf(ut)
	return reflect.MakeSlice(st, 0, 0)
}

// TypeOf returns the type of i
func TypeOf(i interface{}) reflect.Type {
	typ := ValueOf(i).Type()
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ
}

// ValueOf returns the value of i
func ValueOf(i interface{}) reflect.Value {
	val := reflect.ValueOf(i)
	return reflect.Indirect(val)
}

// Foreach applies the given function to each element of i, iff i is a slice
func Foreach(i interface{}, fn func(int, interface{}) bool) {
	val := ValueOf(i)
	typ := val.Type()

	if typ.Kind() == reflect.Slice {
		n := val.Len()
		for i := 0; i < n; i++ {
			if !fn(i, val.Index(i).Interface()) {
				break
			}
		}
	}
}

// Field returns the named field within the given argument.
// The argument should be a struct or *struct.
// All errors are silently reported by returning a zero value.
func Field(i interface{}, name string) reflect.Value /*, error*/ {
	if !IsStructOrStructPtr(i) {
		return reflect.Zero(TypeOf(i)) //, errNotAStruct
	}

	field := ValueOf(i).FieldByName(name)
	if !field.IsValid() {
		return reflect.Zero(TypeOf(i)) //, errUnknownField
	}

	return field //, nil
}

// Zero returns the zero value corresponding to the type of the given parameter.
func Zero(i interface{}) interface{} {
	return reflect.Zero(TypeOf(i))
}

// should be AppendAsSlice - hide for now
// AppendSlice elements of the slice represented by v to the slice represented by ret
// and return the result.  Both parameters must have the same underlying type.
// func AppendSlice(ret, v interface{}) reflect.Value {
//	return reflect.AppendSlice(ValueOf(ret), ValueOf(v))
// }

// not working - hide for now
// func Copy(dst, src interface{}) interface{} {
//	fmt.Printf("Value src: %#v (%T)\n", src, src)
//	fmt.Printf("Value dst: %#v (%T)\n", dst, dst)
//	n := reflect.Copy(ValueOf(dst), ValueOf(src))
//	fmt.Printf("  %v values copied\n", n)
//	fmt.Printf("Value src: %#v (%T)\n", src, src)
//	fmt.Printf("Value dst: %#v (%T)\n", dst, dst)
//	return dst
// }


// debug statements stash
// fmt.Printf("Value i: %#v (%T)\n", i, i)
// fmt.Printf("Value x: %#v (%T)\n", x, x)

