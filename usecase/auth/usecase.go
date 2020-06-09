package auth

import (
	"context"

	"github.com/wahyurudiyan/authentication/repository/auth"
)

type usecase struct {
	auth.Repository
}

type Usecase interface {
}

func NewUsecase(repo auth.Repository) Usecase {
	return &usecase{repo}
}

func (u *usecase) LoginHistory(ctx context.Context, v *entity.LoginHistory)
