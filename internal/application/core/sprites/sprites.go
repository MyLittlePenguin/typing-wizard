package sprites

import raylib "github.com/gen2brain/raylib-go/raylib"

var Deamon raylib.Texture2D
var Mage raylib.Texture2D
var Background raylib.Texture2D

func LoadSprites() {
	Deamon = *load("assets/images/devil_small_purple.png", 64, 61)
	Mage = *load("assets/images/mage.png", 98, 128)
  Background = *load("assets/images/full-moon.jpg", 1280, 854)
}

func load(path string, width, height int32) *raylib.Texture2D {
	img := raylib.LoadImage(path)
	raylib.ImageResize(img, width, height)
	texture := raylib.LoadTextureFromImage(img)
	raylib.UnloadImage(img)
	return &texture
}
