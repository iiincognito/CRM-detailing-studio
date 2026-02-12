package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/iiincognito/CRM-detailing-studio/backend/internal/user/dto"
)

type ServInterface interface {
	Created(ctx context.Context, user *dto.UserDTO) (uuid.UUID, error)
	GetById(ctx context.Context, userId uuid.UUID) (*dto.UserDTO, error)
	GetByEmail(ctx context.Context, userEmail string) (*dto.UserDTO, error)
	Update(ctx context.Context, user *dto.UserDTO) (uuid.UUID, error)
	Delete(ctx context.Context, userId uuid.UUID) error
}
