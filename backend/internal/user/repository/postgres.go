package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/iiincognito/CRM-detailing-studio/backend/internal/user/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	db *pgxpool.Pool
}

func NewPostgres(db *pgxpool.Pool) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) Create(ctx context.Context, user *service.User) (uuid.UUID, error) {
	passHash, err := HashFunc(user.Pass)
	if err != nil {
		return uuid.UUID{}, err
	}
	usrMdl := UserModel{
		ID:       uuid.New(),
		FIO:      user.FIO,
		Email:    user.Email,
		Phone:    user.Phone,
		birthday: user.Birthday,
		passHash: passHash,
	}
	usr := service.User{}
	err = p.db.QueryRow(ctx, "INSERT INTO users ( id,fio, email, phone, birthday, pass_hash) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id; ",
		usrMdl.ID, usrMdl.FIO, usrMdl.Email, usrMdl.Phone, usrMdl.birthday, usrMdl.passHash).Scan(&usr.Id)
	if err != nil {
		return uuid.UUID{}, err
	}
	return usr.Id, nil
}

/*
func (p *Postgres) GetByEmail(ctx context.Context, email string) (*service.User, error) {

}

func (p *Postgres) GetByID(ctx context.Context, id uuid.UUID) (*service.User, error) {

}

func (p *Postgres) Update(ctx context.Context, user *service.User) error {

}

func (p *Postgres) Delete(ctx context.Context, id uuid.UUID) error {

}
*/
