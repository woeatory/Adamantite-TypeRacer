package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user_scores/models"
)

type MockScoreRecord struct {
	mock.Mock
}

func (m *MockScoreRecord) NewScoreRecord(userID string, wpm, accuracy, typos int) error {
	args := m.Called(userID, wpm, accuracy, typos)
	return args.Error(0)
}

func (m *MockScoreRecord) GetAllUsersScoreRecords(userID string) ([]models.ScoreRecord, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.ScoreRecord), args.Error(1)
}

func (m *MockScoreRecord) DeleteScoreRecord(userID string, recordID int) error {
	args := m.Called(userID, recordID)
	return args.Error(0)
}
