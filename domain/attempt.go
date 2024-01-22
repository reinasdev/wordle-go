package domain

import (
	"encoding/json"
	"errors"
	"time"
	"wordle/infra"
	"wordle/utils"
)

// UserAttempt represents a user attempt
type UserAttempt struct {
	IP      *string
	Attempt []CheckWord
	Success *bool
	Date    *time.Time
}

// ValidateAttempt validates a user attempt
func ValidateAttempt(ip *string, attempt []CheckWord, success *bool) (err error) {
	var (
		now      = time.Now()
		attempts = GetUserAttempts(ip, &now)
	)

	if len(attempts) >= 6 {
		return errors.New("try again tomorrow")
	}

	for _, attempt := range attempts {
		if *attempt.Success {
			return errors.New("you already won")
		}
	}

	if err = registerAttempt(ip, attempt, success); err != nil {
		return err
	}

	return nil
}

// registerAttempt registers a user attempt
func registerAttempt(ip *string, attempt []CheckWord, success *bool) (err error) {
	data, err := json.Marshal(attempt)
	if err != nil {
		return err
	}

	if err := infra.RegisterAttempt(ip, data, success, utils.GetPointer(time.Now())); err != nil {
		return nil
	}

	return nil
}

// GetUserAttempts returns all user attempts
func GetUserAttempts(ip *string, date *time.Time) (attempts []UserAttempt) {
	data, err := infra.GetUserAttempts(ip, date)
	if err != nil {
		return nil
	}

	for _, attempt := range data {
		var checkWord []CheckWord
		if err := json.Unmarshal(attempt.Attempt, &checkWord); err != nil {
			return nil
		}

		attempts = append(attempts, UserAttempt{
			IP:      attempt.IP,
			Attempt: checkWord,
			Success: attempt.Success,
			Date:    attempt.Date,
		})
	}

	return attempts
}
