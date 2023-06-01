package models

import (
	"github.com/google/uuid"
	"time"
)

type ScoreRecord struct {
	RecordID  int
	UserID    uuid.UUID
	WPM       int
	Accuracy  int
	Typos     int
	CreatedAt time.Time
}
