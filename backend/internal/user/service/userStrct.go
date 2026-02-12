package service

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id       uuid.UUID
	FIO      string
	Email    string
	Phone    string
	Birthday time.Time
	Pass     string
	Created  time.Time
	Updated  time.Time
}
