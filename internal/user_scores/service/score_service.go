package service

import "github.com/woeatory/Adamantite-TypeRacer/internal/repository"

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

func (scoreService *ScoreService) DeleteScoreRecord(userID string, recordID int) error {
	query := "DELETE FROM scores WHERE record_id = $1 AND user_id = $2"
	stmt, err := scoreService.repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(recordID, userID)
	if err != nil {
		return err
	}
	return nil
}
