package raylib

/*
#cgo pkg-config: raylib
#include "raylib.h"
*/
import "C"
import (
	"image/color"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

// window stuff
func InitWindow(width, height int32, title string) {
	C.InitWindow(cint(width), cint(height), cstring(title))
}

func WindowShouldClose() bool { return bool(C.WindowShouldClose()) }
func CloseWindow()            { C.CloseWindow() }
func BeginDrawing()           { C.BeginDrawing() }
func EndDrawing()             { C.EndDrawing() }

func ClearBackground(color *color.RGBA) { C.ClearBackground(*ccolor(color)) }

// timing stuff
func SetTargetFPS(fps int32) { C.SetTargetFPS(cint(fps)) }
func GetFrameTime() float32  { return float32(C.GetFrameTime()) }
func GetTime() float64       { return float64(C.GetTime()) }

// Text stuff
func DrawFPS(x, y int32) { C.DrawFPS(cint(x), cint(y)) }
func DrawText(text string, x, y int32, fontSize int32, color *color.RGBA) {
	C.DrawText(cstring(text), cint(x), cint(y), cint(fontSize), *ccolor(color))
}

// shape stuff
func DrawRectangle(posX, posY, width, height int32, color *color.RGBA) {
	C.DrawRectangle(cint(posX), cint(posY), cint(width), cint(height), *ccolor(color))
}

func DrawRectangleV(position, size Vector2, color *color.RGBA) {
	C.DrawRectangleV(*position.toCPtr(), *size.toCPtr(), *ccolor(color))
}

func DrawRectangleRec(rec Rectangle, color *color.RGBA) {
  C.DrawRectangleRec(*rec.toCPtr(), *ccolor(color))
}

// keyboard input stuff
func IsKeyPressed(key int) bool       { return bool(C.IsKeyPressed(cint(int32(key)))) }
func IsKeyPressedRepeat(key int) bool { return bool(C.IsKeyPressedRepeat(cint(int32(key)))) }
func IsKeyDown(key int) bool          { return bool(C.IsKeyDown(cint(int32(key)))) }
func IsKeyReleased(key int) bool      { return bool(C.IsKeyReleased(cint(int32(key)))) }
func IsKeyUp(key int) bool            { return bool(C.IsKeyUp(cint(int32(key)))) }
func GetKeyPressed() int              { return int(C.GetKeyPressed()) }
func GetCharPressed() int             { return int(C.GetCharPressed()) }
func SetExitKey(key int)              { C.SetExitKey(cint(int32(key))) }

// other stuff
func GetRandomValue(min, max int32) int32 { return int32(C.GetRandomValue(cint(min), cint(max))) }
