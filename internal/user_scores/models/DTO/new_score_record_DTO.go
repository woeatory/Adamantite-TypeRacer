package DTO

type NewScoreRecordDTO struct {
	WPM      int `json:"WPM" binding:"required"`
	Accuracy int `json:"Accuracy" binding:"required"`
	Typos    int `json:"Typos" binding:"required"`
}
