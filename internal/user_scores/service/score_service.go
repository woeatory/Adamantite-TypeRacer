package service

import (
	"errors"
	"github.com/woeatory/Adamantite-TypeRacer/internal/repository"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user_scores/models"
)

type ScoreRecorder interface {
	NewScoreRecord(userID string, wpm, accuracy, typos int) error
	GetAllUsersScoreRecords(userID string) ([]models.ScoreRecord, error)
	DeleteScoreRecord(userID string, recordID int) error
}

type ScoreService struct {
	repo *repository.Repo
}

func NewScoreService(repo *repository.Repo) *ScoreService {
	return &ScoreService{repo: repo}
}

func (scoreService *ScoreService) NewScoreRecord(userID string, wpm, accuracy, typos int) error {
	query := "INSERT INTO scores (user_id, WPM, accuracy, typos) VALUES ($1, $2, $3, $4)"
	stmt, err := scoreService.repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, wpm, accuracy, typos)
	if err != nil {
		return err
	}
	return nil
}

func (scoreService *ScoreService) GetAllUsersScoreRecords(userID string) ([]models.ScoreRecord, error) {
	query := "SELECT * FROM scores WHERE user_id = $1"
	stmt, err := scoreService.repo.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	var records []models.ScoreRecord
	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var record models.ScoreRecord
		if err := rows.Scan(
			&record.RecordID, &record.UserID, &record.WPM, &record.Accuracy, &record.Typos, &record.CreatedAt,
		); err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func (scoreService *ScoreService) DeleteScoreRecord(userID string, recordID int) error {
	query := "DELETE FROM scores WHERE record_id = $1 AND user_id = $2"
	stmt, err := scoreService.repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(recordID, userID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("nothing was deleted")
	}
	return nil
}
