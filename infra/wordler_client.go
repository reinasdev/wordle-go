package infra

import (
	"encoding/json"
	"net/http"
	"time"
	"wordle/config"
)

// WordleResponse represents the response from the wordle API
type WordleResponse struct {
	Solution *string `json:"solution"`
}

// GetDailyWord returns the daily word
func GetDailyWord() (dailyWord *string, err error) {
	var (
		today          = time.Now().Format("2006-01-02")
		wordleResponse WordleResponse
		cfg            = config.GetConfig()
	)

	response, err := http.Get(cfg.WordleURL + today + ".json")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&wordleResponse); err != nil {
		return nil, err
	}

	return wordleResponse.Solution, nil
}
