package dto

import (
	"github.com/google/uuid"
	"time"
)

type UserDTO struct {
	Id       uuid.UUID `json:"id"`
	FIO      string    `json:"fio"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Birthday time.Time `json:"birthday"`
	Pass     string    `json:"pass"`
	Created  time.Time
	Updated  time.Time
}
