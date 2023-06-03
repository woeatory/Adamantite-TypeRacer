package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user_scores/models/DTO"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user_scores/service"
	"net/http"
)

type ScoreController struct {
	scoreService service.ScoreRecorder
}

func NewScoreController(service service.ScoreRecorder) *ScoreController {
	return &ScoreController{scoreService: service}
}

func (scoreController *ScoreController) AddNewScoreRecord(c *gin.Context) {
	var input DTO.NewScoreRecordDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)
	userID := session.Get("id").(string)
	err := scoreController.scoreService.NewScoreRecord(userID, input.WPM, input.Accuracy, input.Typos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while inserting new record:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record Added Successfully"})
}

func (scoreController *ScoreController) GetAllUsersRecords(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("id").(string)
	res, err := scoreController.scoreService.GetAllUsersScoreRecords(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error getting users records:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": res})
}

func (scoreController *ScoreController) DeleteScoreRecord(c *gin.Context) {
	var input DTO.DeleteScoreRecord
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)
	userID := session.Get("id").(string)
	err := scoreController.scoreService.DeleteScoreRecord(userID, input.RecordID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "Record Deleted Successfully"})
}
