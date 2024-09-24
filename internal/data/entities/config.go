package entities

import (
	"fmt"
	"type_wizard/internal/data/model"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Config struct {
	EnchantmentReleaseKey int
	Highscore             int
}

func NewConfig() Config {
	return Config{
		EnchantmentReleaseKey: raylib.KeySpace,
		Highscore:             0,
	}
}

func (self Config) ToModel() model.Config {
	configModel := model.Config{
		Highscore: self.Highscore,
	}
	switch self.EnchantmentReleaseKey {
	case raylib.KeyEnter:
		configModel.EnchantmentReleaseKey = "ENTER"
	default:
		configModel.EnchantmentReleaseKey = "SPACE"
	}
	return configModel
}

func ConfigFromModel(configModel model.Config) Config {
	config := NewConfig()
	switch configModel.EnchantmentReleaseKey {
	case "ENTER":
		config.EnchantmentReleaseKey = raylib.KeyEnter
	case "SPACE":
		config.EnchantmentReleaseKey = raylib.KeySpace
	default:
		config.EnchantmentReleaseKey = raylib.KeySpace
	}
	config.Highscore = configModel.Highscore
	return config
}

func (self Config) ToString() string {
  model := self.ToModel()
  return fmt.Sprintf("Config(EnchantmentReleaseKey: %s, Highscore: %d", model.EnchantmentReleaseKey, model.Highscore)
}
