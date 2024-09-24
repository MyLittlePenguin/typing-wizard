package highscore

import (
	"fmt"
	viewmanager "type_wizard/internal/application/core/view_manager"
	"type_wizard/internal/application/core/window"
	configuc "type_wizard/internal/domain/usecase/config_uc"

	"github.com/gen2brain/raylib-go/raygui"
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Highscore struct {
}

var (
	titleX = int32(0)
	btnPos = raylib.Rectangle{
		X:      float32(window.Width-200) / 2,
		Y:      300,
		Width:  200,
		Height: 30,
	}
)

func New() viewmanager.View {
	titleX = int32(window.Width-raylib.MeasureText("HighScore", 30)) / 2
	return Highscore{}
}

func (self Highscore) Update() {}

func (self Highscore) View() {
	raylib.DrawText("HighScore", titleX, 200, 30, raylib.RayWhite)
	score := fmt.Sprintf("%d", configuc.ReadConfig().Highscore)
	scoreX := (window.Width - raylib.MeasureText(score, 25)) / 2
	raylib.DrawText(score, scoreX, 250, 25, raylib.RayWhite)
	clicked := raygui.Button(
		btnPos,
		"Back",
	)
	if clicked {
		viewmanager.Open("menu")
	}
}
