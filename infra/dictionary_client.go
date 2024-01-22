package infra

import (
	"errors"
	"net/http"
	"wordle/config"
)

// CheckIfWordExist checks if a word exists in the dictionary
func CheckIfWordExist(word *string) (err error) {
	cfg := config.GetConfig()

	response, err := http.Get(cfg.DictionaryURL + *word)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return errors.New("word not found")
	}

	return nil
}
