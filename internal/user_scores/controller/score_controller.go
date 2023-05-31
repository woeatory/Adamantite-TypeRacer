package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user_scores/models/DTO"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user_scores/service"
	"net/http"
)

type ScoreController struct {
	scoreService *service.ScoreService
}

func NewScoreController(service *service.ScoreService) *ScoreController {
	return &ScoreController{scoreService: service}
}

func (scoreController *ScoreController) AddNewScore(c *gin.Context) {
	var input DTO.NewScoreRecordDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)
	var userID string
	v := session.Get("id")
	if v == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	} else {
		userID = v.(string)
	}
	err := scoreController.scoreService.NewScoreRecord(userID, input.WPM, input.Accuracy, input.Typos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error while inserting new record"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "record added successfully"})
}

func (scoreController *ScoreController) DeleteScoreRecord(c *gin.Context) {
	var input DTO.DeleteScoreRecord
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)
	var userID string
	v := session.Get("id")
	if v == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	} else {
		userID = v.(string)
	}
	err := scoreController.scoreService.DeleteScoreRecord(userID, input.RecordID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "record deleted successfully"})
}
