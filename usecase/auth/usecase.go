package auth

import (
	"context"

	authEntity "github.com/wahyurudiyan/authentication/entity/auth"
	"github.com/wahyurudiyan/authentication/repository/account"
	"github.com/wahyurudiyan/authentication/repository/auth"
)

type usecase struct {
	account account.Repository
	auth    auth.Repository
}

type Usecase interface {
	Login(ctx context.Context, v *authEntity.Login) error
	LoginHistory(ctx context.Context, v *authEntity.LoginHistory) error
}

func NewUsecase(accountRepo account.Repository, authRepo auth.Repository) Usecase {
	return &usecase{accountRepo, authRepo}
}

func (u *usecase) Login(ctx context.Context, v *authEntity.Login) error {
	return nil
}

func (u *usecase) LoginHistory(ctx context.Context, v *authEntity.LoginHistory) error {
	return nil
}
