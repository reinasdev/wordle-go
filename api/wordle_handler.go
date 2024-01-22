package api

import (
	"github.com/gin-gonic/gin"
	"time"
	"wordle/utils"

	"net/http"
	"wordle/domain"
	"wordle/infra"
)

// RegisterHandlers registers the handlers
func RegisterHandlers(r *gin.RouterGroup) {
	r.GET("/check", CheckHandler)
	r.GET("/attempts/:date", AttemptsHandler)
}

// CheckHandler godoc
// @Summary Check if a word is correct
// @Description Check if a word is correct
func CheckHandler(c *gin.Context) {
	word := c.Query("word")

	if err := infra.CheckIfWordExist(&word); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attempt, success, err := domain.VerifyWord(&word)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = domain.ValidateAttempt(utils.GetPointer(c.ClientIP()), attempt, success); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"check": attempt, "success": success})
}

// AttemptsHandler godoc
// @Summary Get user attempts
// @Description Get user attempts
func AttemptsHandler(c *gin.Context) {
	date, err := time.Parse("2006-01-02", c.Param("date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"attempts": domain.GetUserAttempts(utils.GetPointer(c.ClientIP()), &date)})
}
