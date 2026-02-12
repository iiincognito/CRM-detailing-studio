package main

import (
	"context"
	"fmt"
	"github.com/iiincognito/CRM-detailing-studio/backend/internal/db"
	"github.com/iiincognito/CRM-detailing-studio/backend/internal/user/dto"
	"github.com/iiincognito/CRM-detailing-studio/backend/internal/user/repository"
	"time"
)

func main() {
	bgctx := context.Background()
	ctx, cancel := context.WithCancel(bgctx)
	db, err := db.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	defer cancel()
	tmp := time.Date(2002, 07, 28, 0, 0, 0, 0, time.UTC)
	tmp.Format("2006-01-02")

	usr := dto.UserDTO{
		FIO:      "Дровосек Айдар Динр",
		Email:    "drova228@mail.ru",
		Phone:    "89991639447",
		Birthday: tmp,
		Pass:     "almazik5560839",
	}
	postgres := repository.NewPostgres(db)
	id, err := postgres.Created(ctx, &usr)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
