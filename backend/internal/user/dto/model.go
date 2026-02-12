package dto

import (
	"github.com/google/uuid"
	"time"
)

type UserModel struct {
	ID       uuid.UUID
	FIO      string
	Email    string
	Phone    string
	Birthday time.Time
	PassHash string
	Created  time.Time
	Updated  time.Time
}
