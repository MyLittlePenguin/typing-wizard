package raylib

/*
#cgo pkg-config: raylib
#include "raylib.h"
*/
import "C"
import (
	"image/color"
	"unsafe"
)

func ccolor(color *color.RGBA) *C.Color { return (*C.Color)(unsafe.Pointer(color)) }
func cint(i int32) C.int { return (C.int)(i) }
func cstring(str string) *C.char { return C.CString(str) }
