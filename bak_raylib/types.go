package raylib

/*
#cgo pkg-config: raylib
#include "raylib.h"
*/
import "C"
import "unsafe"

type Vector2 struct {
	X, Y float32
}

func (self Vector2) toCPtr() *C.Vector2 {
	return (*C.Vector2)(unsafe.Pointer(&self))
}

func (self Vector2) Add(other Vector2) Vector2 {
	return Vector2{self.X + other.X, self.Y + other.Y}
}

func (self Vector2) Substract(other Vector2) Vector2 {
	return Vector2{self.X - other.X, self.Y - other.Y}
}

func (self Vector2) Scale(factor float32) Vector2 {
  return Vector2{self.X * factor, self.Y * factor}
}

type Rectangle struct {
  X, Y, Width, Height float32
}

func (self Rectangle) toCPtr() *C.Rectangle {
  return (*C.Rectangle)(unsafe.Pointer(&self))
}
