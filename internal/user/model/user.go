package model

import (
	"crypto"
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserID       uuid.UUID
	Username     string
	PasswordHash crypto.Hash
	CreatedAt    time.Time
}
