package service

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	id        uuid.UUID
	email     string
	passHash  string
	name      string
	age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
