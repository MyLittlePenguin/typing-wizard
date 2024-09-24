package fonts

import raylib "github.com/gen2brain/raylib-go/raylib"

const (
	FontSize = 20
	FontSpacing  = 1
)

var currentFont raylib.Font

var roboto raylib.Font
func Roboto() raylib.Font {
  return roboto
}

func LoadFonts() {
  roboto = raylib.LoadFontEx("assets/fonts/roboto/Roboto-Medium.ttf", 20, nil, 250)
}

func SelectFont(font raylib.Font) {
  currentFont = font
}

func Font() raylib.Font {
  return currentFont
}
