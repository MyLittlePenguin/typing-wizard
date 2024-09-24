package repo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"type_wizard/internal/data/entities"
	"type_wizard/internal/data/model"
)

const (
	configFileName = "typing-wizard.json"
)

func GetCommonDictionary() []string {
	top1000, err := os.Open("assets/dictionaries/top1000.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(top1000)
	dict := make([]string, 0, 1000)
	for scanner.Scan() {
		dict = append(dict, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	top1000.Close()
	return dict
}

func ReadConfig() entities.Config {
	content, err := os.ReadFile(configFileName)
	config := entities.NewConfig()
	if err != nil {
		return config
	}
  fmt.Printf("read config: %s\n", string(content))
	configModel := model.Config{}
	err = json.Unmarshal(content, &configModel)
	if err != nil {
		return config
	}

  config = entities.ConfigFromModel(configModel)
	return config
}

func WriteConfig(config entities.Config) {
	json, _ := json.Marshal(config.ToModel())
	os.WriteFile(configFileName, json, 0744)
}
