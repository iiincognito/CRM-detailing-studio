package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/iiincognito/CRM-detailing-studio/backend/internal/user/dto"
	"github.com/iiincognito/CRM-detailing-studio/backend/internal/user/repository"
)

type Ser struct {
	rpstr repository.UserRpstr
}

func NewSer(rpstr repository.UserRpstr) ServInterface {
	return &Ser{
		rpstr: rpstr,
	}
}

func (s *Ser) Created(ctx context.Context, user *dto.UserDTO) (uuid.UUID, error) {
	
}

func (s *Ser) GetById(ctx context.Context, userId uuid.UUID) (*dto.UserDTO, error) {

}

func (s *Ser) GetByEmail(ctx context.Context, userEmail string) (*dto.UserDTO, error) {

}

func (s *Ser) Update(ctx context.Context, user *dto.UserDTO) (uuid.UUID, error) {

}

func (s *Ser) Delete(ctx context.Context, userId uuid.UUID) error {

}
