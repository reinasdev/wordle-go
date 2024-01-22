package migrations

import (
	"log"
	"os"
	"wordle/config"
	"wordle/database"
)

// SetupDatabase sets up the database
func SetupDatabase() (err error) {
	cfg := config.GetConfig()

	if cfg.DatabasePath == "" {
		return nil
	}

	_, err = os.Create(cfg.DatabaseFile)
	if err != nil {
		return err
	}

	db, err := database.LoadDatabase()
	if err != nil {
		return err
	}
	defer db.DB.Close()

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	path = path + "/migrations/database.sql"

	raw, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.DB.Exec(string(raw))
	if err != nil {
		return err
	}

	return nil
}
