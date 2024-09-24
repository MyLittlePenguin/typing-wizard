package dictionaryuc

import (
	// "type_wizard/internal/application/core/raylib"
  raylib "github.com/gen2brain/raylib-go/raylib"
	"type_wizard/internal/data/repo"
)

var dictionary = make([]string, 0, 1000)
var difficulty = 0.0

func GetDictionary() []string {
	if len(dictionary) == 0 {
		commons := repo.GetCommonDictionary()
		dictionary = append(dictionary, commons...)
	}
	return dictionary
}

func GetWord() string {
	min := int(difficulty * 200)
	max := int(difficulty + 1) * 200
	dict := GetDictionary()
	length := len(dict)

	if length == 0 {
		return "Crapple"
	}

	if min >= length-200 {
    min = length - 200
	}
  if max > length {
    max = length
  }

  dict = dict[min: max]
  length = len(dict)
	i := raylib.GetRandomValue(0, int32(length-1))
	return (dict)[i]
}

func SetDifficulty(value int) {
	difficulty = float64(value) * 0.5
}

func Reset() {
	difficulty = 0
}
