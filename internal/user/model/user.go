package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserID       uuid.UUID
	Username     string
	PasswordHash string
	CreatedAt    time.Time
}
