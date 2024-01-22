package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"wordle/api"
	"wordle/config"
	"wordle/migrations"
)

// @title Wordle API
// @description This is a wordle API server.
func main() {
	config.LoadConfig()

	var (
		r   = gin.Default()
		cfg = config.GetConfig()
		g   = r.Group("/api")
	)

	if err := migrations.SetupDatabase(); err != nil {
		log.Fatal(err)
		return
	}

	api.RegisterHandlers(g)

	if err := r.Run(":" + strconv.Itoa(cfg.Port)); err != nil {
		log.Fatal(err)
		return
	}
}
