package deamonuc

import (
	// "type_wizard/internal/application/core/raylib"
	"type_wizard/internal/application/core/fonts"
	"type_wizard/internal/data/entities"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

func NewDeamon(enchantment string, winWidth int32) entities.Deamon {
	size := raylib.MeasureTextEx(fonts.Font(), enchantment, fonts.FontSize, fonts.FontSpacing)
	return entities.Deamon{
		Pos: raylib.Vector2{
			X: float32(raylib.GetRandomValue(0, winWidth-100)),
			Y: 15,
		},
		Size: raylib.Vector2{
			X: size.X + 15,
			Y: size.Y + 10,
		},
		Enchantment: enchantment,
	}
}

func Move(deamon *entities.Deamon, movement raylib.Vector2) {
	deamon.Pos = raylib.Vector2{
		X: deamon.Pos.X + movement.X,
		Y: deamon.Pos.Y + movement.Y,
	}
}

func Fall(deamon *entities.Deamon, time float32) {
	Move(deamon, raylib.Vector2{X: 0, Y: time * 25})
}

func Chant(deamon *entities.Deamon, enchantment string) bool {
	return deamon.Enchantment == enchantment
}
