package main

import (
	"image/color"
	"type_wizard/internal/application/core/fonts"
	"type_wizard/internal/application/core/sprites"
	viewmanager "type_wizard/internal/application/core/view_manager"
	"type_wizard/internal/application/core/window"
	"type_wizard/internal/application/game"
	"type_wizard/internal/application/highscore"
	"type_wizard/internal/application/menu"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	raylib.InitWindow(window.Width, window.Height, window.Title)
	// raylib.SetTargetFPS(FPS)
	fonts.LoadFonts()
	fonts.SelectFont(fonts.Roboto())
	registerViews()
	viewmanager.Open("menu")
	sprites.LoadSprites()

	for !raylib.WindowShouldClose() {
		raylib.BeginDrawing()
		raylib.ClearBackground(color.RGBA{90, 90, 100, 255})
		raylib.SetExitKey(raylib.KeyNull)

		view := viewmanager.Get()
		view.Update()
		view.View()

		raylib.DrawFPS(1100, 10)
		raylib.EndDrawing()
	}
	raylib.CloseWindow()
}

func registerViews() {
	viewmanager.Register("game", game.New)
	viewmanager.Register("menu", menu.New)
	viewmanager.Register("highscore", highscore.New)
}
