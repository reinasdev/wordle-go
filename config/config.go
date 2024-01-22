package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

var config *Config

// LoadConfig tries to read from `wordle_CONFIG` file a valid JSON with all settings
func LoadConfig() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	path = path + "/config.json"

	if val, set := os.LookupEnv("WORDLE_CONFIG"); set && val != "" {
		path = val
	} else {
		log.Println("`WORDLE_CONFIG` not set or empty, using default path: ", path)
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(raw, &config); err != nil {
		log.Fatal(err)
	}

	if err := validateConfig(); err != nil {
		log.Fatal(err)
	}
}

// GetConfig returns a pointer to a Config struct which holds a valid config
func GetConfig() *Config {
	if config == nil {
		log.Fatal("config was not successfully loaded")
	}
	return config
}

// validateConfig validates the config
func validateConfig() error {
	if config == nil {
		return errors.New("config was not successfully loaded")
	}

	if config.WordleURL == "" {
		config.WordleURL = "https://www.nytimes.com/svc/wordle/v2/"
	}

	if config.Port == 0 {
		config.Port = 8080
	}

	if config.DictionaryURL == "" {
		config.DictionaryURL = "https://api.dictionaryapi.dev/api/v2/entries/en/"
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	if config.DatabaseFile == "" {
		config.DatabaseFile = "database.db"
	}

	if config.DatabasePath == "" {
		config.DatabasePath = path + "/" + config.DatabaseFile
	} else {
		config.DatabasePath = config.DatabasePath + "/" + config.DatabaseFile
	}

	return nil
}
