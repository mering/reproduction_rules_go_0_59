// Package world provides a go wrapper for creating and deleting world objects from go.
package world

import "C"

import (
	"unsafe"
)

type World struct {
	impl unsafe.Pointer
}
