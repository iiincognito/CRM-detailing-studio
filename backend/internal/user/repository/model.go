package repository

import (
	"github.com/google/uuid"
	"time"
)

type UserModel struct {
	ID       uuid.UUID
	FIO      string
	Email    string
	Phone    string
	birthday time.Time
	passHash string
}
