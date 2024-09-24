package menu

import (
	"os"
	viewmanager "type_wizard/internal/application/core/view_manager"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Menu struct{}

func New() viewmanager.View {
	return &Menu{}
}

func (self Menu) Update() {
}

var buttons = []Button{
	Button{
		Nr:   0,
		Text: "Play",
		Action: func() {
			viewmanager.Open("game")
		},
	},
	Button{
		Nr:   1,
		Text: "Highscore",
		Action: func() {
			viewmanager.Open("highscore")
		},
	},
	Button{
		Nr:   2,
		Text: "Quit",
		Action: func() {
			raylib.CloseWindow()
			os.Exit(0)
		},
	},
}

func (self Menu) View() {
	for _, button := range buttons {
		button.Render()
	}
}
