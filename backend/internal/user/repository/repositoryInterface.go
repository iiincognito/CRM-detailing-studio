package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/iiincognito/CRM-detailing-studio/backend/internal/user/dto"
)

type UserRpstr interface {
	Created(ctx context.Context, user *dto.UserDTO) (uuid.UUID, error)    // Create — создаёт нового пользователя в БД.
	GetByEmail(ctx context.Context, email string) (*dto.UserModel, error) // GetByEmail — ищет пользователя по email.
	GetByID(ctx context.Context, id uuid.UUID) (*dto.UserModel, error)    // GetByID — ищет пользователя по ID.
	Update(ctx context.Context, user *dto.UserDTO) error                  // Update — обновляет существующие поля пользователя (name, age и т.д.).
	Delete(ctx context.Context, id uuid.UUID) error
}
