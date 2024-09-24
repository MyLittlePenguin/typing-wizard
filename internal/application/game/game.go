package game

import (
	"fmt"
	"image/color"

	// "type_wizard/internal/application/core/raylib"
	"type_wizard/internal/application/core/fonts"
	"type_wizard/internal/application/core/sprites"
	viewmanager "type_wizard/internal/application/core/view_manager"
	"type_wizard/internal/application/core/window"
	"type_wizard/internal/application/deamons"
	"type_wizard/internal/data/entities"
	configuc "type_wizard/internal/domain/usecase/config_uc"
	deamonuc "type_wizard/internal/domain/usecase/deamon_uc"
	dictionaryuc "type_wizard/internal/domain/usecase/dictionary_uc"

	"github.com/gen2brain/raylib-go/raygui"
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	startTime                      float64
	secsSinceDeamon, secsPerDeamon float32
	score                          int
	deamons                        []entities.Deamon
	lost                           bool
	enchantment                    string
	time                           float64
	enchantmentBox                 raylib.Rectangle
	enchantmentReleaseKey          int32
}

const (
	DEAMON_SPAWN_TIME float32 = 6
)

func New() viewmanager.View {
	return &Game{
		startTime:       raylib.GetTime(),
		score:           0,
		secsPerDeamon:   DEAMON_SPAWN_TIME,
		secsSinceDeamon: 15,
		deamons:         make([]entities.Deamon, 0, 10),
		lost:            false,
		enchantment:     "",
		enchantmentBox: raylib.Rectangle{
			X:      0,
			Y:      float32(window.Height - 50),
			Width:  window.Width,
			Height: 50,
		},
		enchantmentReleaseKey: int32(configuc.ReadConfig().EnchantmentReleaseKey),
	}
}

func (self *Game) Update() {
	if self.lost {
		return
	}
	self.time = raylib.GetTime() - self.startTime
	if self.secsPerDeamon < self.secsSinceDeamon {
		self.CreateDeamon()
		self.secsSinceDeamon = 0
	} else {
		self.secsSinceDeamon += raylib.GetFrameTime()
	}
	self.MoveDeamons()
	self.ReadEnchantment()
	self.CheckDeamons()
}

func (self *Game) CreateDeamon() {
	newEnchantment := dictionaryuc.GetWord()
	if raylib.GetRandomValue(0, 40) == 1 {
		newEnchantment = "Crapple"
	}
	newDeamon := deamonuc.NewDeamon(newEnchantment, window.Width)
	self.deamons = append(self.deamons, newDeamon)
}

func (self *Game) MoveDeamons() {
	l := len(self.deamons)
	for i := 0; i < l; i++ {
		deamon := &self.deamons[i]
		deamonuc.Fall(deamon, raylib.GetFrameTime())
	}
}

func (self *Game) ReadEnchantment() {
	r := 0
	r = int(raylib.GetCharPressed())
	if r > 0 {
		if string(r) != " " {
			self.enchantment += string(r)
		}
	}

	if raylib.IsKeyPressed(self.enchantmentReleaseKey) {
		match := false
		l := len(self.deamons)
		for i := 0; i < l; i++ {
			deamon := &self.deamons[i]
			if deamonuc.Chant(deamon, self.enchantment) {
				self.removeDeamon(i)
				self.score += len(self.enchantment)
				dictionaryuc.SetDifficulty(self.score / 50)
				// if self.secsPerDeamon > 1 {
          self.secsPerDeamon = (self.secsPerDeamon - 1) * 0.95 + 1
					// self.secsPerDeamon *= 0.95
				// }
				match = true
				break
			}
		}
		if !match {
			self.score -= 2
		}
		self.enchantment = ""
	}
}

func (self *Game) removeDeamon(idx int) {
	if idx >= 0 && idx < len(self.deamons) {
		self.deamons = append(self.deamons[:idx], self.deamons[idx+1:]...)
	}
}

func (self *Game) CheckDeamons() {
	for _, deamon := range self.deamons {
		border := float32(window.Height - 50 - deamon.Size.Y)
		if deamon.Pos.Y >= border {
			self.lost = true
			configuc.UpdateHighscore(self.score)
		}
	}
}

func (self Game) View() {
	raylib.DrawTexture(sprites.Background, 0, 0, raylib.RayWhite)
	for _, deamon := range self.deamons {
		deamons.Render(&deamon)
	}
	score(self.score, window.Width/2-50, 10)
	if self.lost {
		raylib.ClearBackground(color.RGBA{
			raylib.DarkGray.R,
			raylib.DarkGray.G,
			raylib.DarkGray.B,
			200,
		})
		raylib.DrawText("You Lost", window.Width/2-55, window.Height/2-15, 30, raylib.RayWhite)
		clicked := raygui.Button(
			raylib.Rectangle{
				X:      float32(window.Width-200) / 2,
				Y:      window.Height/2 + 30,
				Width:  200,
				Height: 40,
			},
			"Back",
		)
		if clicked {
			viewmanager.Open("menu")
		}
		return
	}
	raylib.DrawText(
		fmt.Sprintf("time: %.2f", self.time), 10, 10, 20, raylib.RayWhite,
	)

  mageOffset := 5 * ((len(self.enchantment) % 2) * -1)
	raylib.DrawTexture(
		sprites.Mage,
		(window.Width-sprites.Mage.Width)/2 + int32(mageOffset),
		window.Height-50-sprites.Mage.Height,
		raylib.RayWhite,
	)
	size := raylib.MeasureTextEx(fonts.Font(), self.enchantment, 20, 1)
	raylib.DrawRectangleRec(self.enchantmentBox, color.RGBA{46, 26, 4, 255})
	raylib.DrawTextEx(
		fonts.Font(),
		self.enchantment,
		raylib.Vector2{
			X: float32((window.Width - size.X) / 2),
			Y: float32(window.Height - 35),
		},
		20,
		1,
		raylib.RayWhite,
	)
}

func score(score int, x, y int32) {
	raylib.DrawText(fmt.Sprintf("Score: %d", score), x, y, 20, raylib.RayWhite)
}
