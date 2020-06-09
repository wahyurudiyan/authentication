package account

import (
	"context"

	"github.com/wahyurudiyan/authentication/entity/account"
	account "github.com/wahyurudiyan/authentication/repository/account"
)

type usecase struct {
	repo account.Repository
}

type Usecase interface {
	CreateAcccount(ctx context.Context, user []*account.Users) error
	GetAccountByID(ctx context.Context, id []string) ([]*account.Users, error)
	GetAccountByUniqueID(ctx context.Context, uid []string) ([]*account.Users, error)
	GetAllAcccount(ctx context.Context) ([]*account.Users, error)
	UpdateAccount(ctx context.Context, uid []string, payload []*account.Users) error
	DeleteAccount(ctx context.Context, uid []string) error
}

func NewAccountUsecase(repo account.Repository) Usecase {
	return &usecase{repo}
}

func (u *usecase) CreateAcccount(ctx context.Context, user []*account.Users) error {
	return u.repo.CreateAcccount(ctx, user)
}

func (u *usecase) GetAccountByID(ctx context.Context, id []string) ([]*account.Users, error) {
	return u.repo.GetAccountByID(ctx, id)
}

func (u *usecase) GetAccountByUniqueID(ctx context.Context, uid []string) ([]*account.Users, error) {
	return u.repo.GetAccountByUniqueID(ctx, uid)
}

func (u *usecase) GetAllAcccount(ctx context.Context) ([]*account.Users, error) {
	return u.repo.GetAllAcccount(ctx)
}

func (u *usecase) UpdateAccount(ctx context.Context, uid []string, payload []*account.Users) error {
	return u.repo.UpdateAccount(ctx, uid, payload)
}

func (u *usecase) DeleteAccount(ctx context.Context, uid []string) error {
	return u.repo.DeleteAccount(ctx, uid)
}
