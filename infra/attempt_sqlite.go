package infra

import (
	"github.com/Masterminds/squirrel"
	"time"
	"wordle/database"
)

// UserAttempt represents a user attempt
type UserAttempt struct {
	ID      *int       `json:"id"`
	IP      *string    `json:"ip"`
	Attempt []byte     `json:"attempts"`
	Success *bool      `json:"success"`
	Date    *time.Time `json:"date"`
}

// RegisterAttempt registers an attempt
func RegisterAttempt(ip *string, attempt []byte, success *bool, date *time.Time) (err error) {
	db, err := database.LoadDatabase()
	if err != nil {
		return err
	}
	defer db.DB.Close()

	query, args, _ := squirrel.
		Insert("attempts").
		Columns("ip", "attempt", "success", "date").
		Values(ip, attempt, success, date).
		ToSql()

	_, err = db.DB.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// GetUserAttempts gets a user's attempts
func GetUserAttempts(ip *string, date *time.Time) (attempts []UserAttempt, err error) {
	db, err := database.LoadDatabase()
	if err != nil {
		return nil, err
	}
	defer db.DB.Close()

	query, args, _ := squirrel.
		Select("*").
		From("attempts").
		Where(squirrel.Eq{
			"ip": ip,
		}).
		Where(squirrel.Like{
			"date": date.Format("2006-01-02") + `%`,
		}).ToSql()

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id      int
			ip      string
			attempt []byte
			success bool
			date    time.Time
		)

		if err := rows.Scan(&id, &ip, &attempt, &success, &date); err != nil {
			return nil, err
		}

		attempts = append(attempts, UserAttempt{
			ID:      &id,
			IP:      &ip,
			Attempt: attempt,
			Success: &success,
			Date:    &date,
		})
	}

	return attempts, nil
}
