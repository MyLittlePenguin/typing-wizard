package deamons

import (
	"type_wizard/internal/application/core/fonts"
	"type_wizard/internal/application/core/sprites"
	"type_wizard/internal/data/entities"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

func Render(deamon *entities.Deamon) {
	enchantmentSize := raylib.MeasureTextEx(fonts.Font(), deamon.Enchantment, fonts.FontSize, fonts.FontSpacing)
	deamonX := deamon.Pos.X + (enchantmentSize.X-64)/2
	raylib.DrawTexture(sprites.Deamon, int32(deamonX), int32(deamon.Pos.Y-61), raylib.White)
	raylib.DrawTextEx(
		fonts.Font(),
		deamon.Enchantment,
		raylib.Vector2{
			X: float32(deamon.Pos.X),
			Y: float32(deamon.Pos.Y) + 5,
		},
		fonts.FontSize,
		fonts.FontSpacing,
		raylib.RayWhite,
	)
}
