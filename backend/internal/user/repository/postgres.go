package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/iiincognito/CRM-detailing-studio/backend/internal/user/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

func NewPostgres(db *pgxpool.Pool) UserRpstr {
	return &postgres{db: db}
}

func (p *postgres) Created(ctx context.Context, user *dto.UserDTO) (uuid.UUID, error) {
	passHash, err := HashFunc(user.Pass)
	if err != nil {
		return uuid.UUID{}, err
	}
	usrMdl := dto.UserModel{
		ID:       uuid.New(),
		FIO:      user.FIO,
		Email:    user.Email,
		Phone:    user.Phone,
		Birthday: user.Birthday,
		PassHash: passHash,
	}
	usr := dto.UserDTO{}
	err = p.db.QueryRow(ctx, "INSERT INTO users ( id,fio, email, phone, birthday, pass_hash) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id; ",
		usrMdl.ID, usrMdl.FIO, usrMdl.Email, usrMdl.Phone, usrMdl.Birthday, usrMdl.PassHash).Scan(&usr.Id)
	if err != nil {
		return uuid.UUID{}, err
	}
	return usr.Id, nil
}

func (p *postgres) GetByEmail(ctx context.Context, email string) (*dto.UserModel, error) {
	model := dto.UserModel{}
	err := p.db.QueryRow(ctx, "SELECT id,fio,email,phone,birthday FROM users WHERE email=$1",
		email).Scan(&model.ID, &model.FIO, &model.Email, &model.Phone, &model.Birthday)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (p *postgres) GetByID(ctx context.Context, id uuid.UUID) (*dto.UserModel, error) {
	model := dto.UserModel{}
	err := p.db.QueryRow(ctx, "SELECT id,fio,email,phone,birthday FROM users WHERE id=$1",
		id).Scan(&model.ID, &model.FIO, &model.Email, &model.Phone, &model.Birthday)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (p *postgres) Update(ctx context.Context, user *dto.UserDTO) error {
	return nil
}

func (p *postgres) Delete(ctx context.Context, id uuid.UUID) error {
	res, err := p.db.Exec(ctx, "DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("user not found")
	}
	return nil
}
