package menu

import (
	"type_wizard/internal/application/core/window"

	"github.com/gen2brain/raylib-go/raygui"
	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	buttonWidth  = 200
	buttonHeight = 40
	buttonX      = float32(window.Width - buttonWidth) / 2
)

type Button struct {
	Nr     int
	Text   string
	Action func()
}

func (self Button) Render() {
	clicked := raygui.Button(
		raylib.Rectangle{
			X:      buttonX,
			Y:      buttonHeight * float32(self.Nr+2) * 1.5,
			Width:  buttonWidth,
			Height: buttonHeight,
		},
		self.Text,
	)
	if clicked {
		self.Action()
	}
}
