package configuc

import (
	"type_wizard/internal/data/entities"
	"type_wizard/internal/data/repo"
)

var config = repo.ReadConfig()

func ReadConfig() entities.Config {
	return config
}

func UpdateHighscore(score int) {
	if config.Highscore < score {
		config.Highscore = score
		repo.WriteConfig(config)
	}
}
