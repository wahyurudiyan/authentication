package account

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/wahyurudiyan/authentication/entity/usersEntity"
)

type repository struct {
	conn *sql.DB
}
type Repository interface {
	CreateAcccount(ctx context.Context, user []*usersEntity.Users) error
	GetAccountByID(ctx context.Context, id []string) ([]*usersEntity.Users, error)
	GetAccountByUniqueID(ctx context.Context, uid []string) ([]*usersEntity.Users, error)
	GetAllAcccount(ctx context.Context) ([]*usersEntity.Users, error)
	UpdateAccount(ctx context.Context, uid []string, payload []*usersEntity.Users) error
	DeleteAccount(ctx context.Context, uid []string) error
}

func NewAccountRepository(mysql *sql.DB) Repository {
	return &repository{mysql}
}

func isDeleted(flag time.Time) bool {
	if flag.IsZero() {
		return false
	}

	return true
}

func (r *repository) CreateAcccount(ctx context.Context, user []*usersEntity.Users) error {
	var args []string
	for _, v := range user {
		arg := fmt.Sprintf("('%v', '%v', '%v', '%v', '%v', '%v', '%v')",
			v.AccountUniqueID, v.Firstname, v.Surename, v.Email, v.Phone, time.Now(), time.Now())
		args = append(args, arg)
	}

	val := strings.Join(args, ", ")
	log.Println(val)
	stmt := fmt.Sprintf("INSERT INTO users (account_unique_id, firstname, surename, email, phone, created_at, updated_at) VALUE %s", val)

	_, err := r.conn.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAccountByID(ctx context.Context, id []string) ([]*usersEntity.Users, error) {
	var results []*usersEntity.Users

	for _, v := range id {
		var row usersEntity.Users
		r := r.conn.QueryRowContext(ctx, "SELECT * FROM users WHERE id=?", v)
		err := r.Scan(
			&row.ID,
			&row.AccountUniqueID,
			&row.Firstname,
			&row.Surename,
			&row.Email,
			&row.Phone,
			&row.CreatedAt,
			&row.UpdatedAt,
			&row.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		if !isDeleted(row.DeletedAt) {
			results = append(results, &row)
		}
	}

	return results, nil
}

func (r *repository) GetAccountByUniqueID(ctx context.Context, uid []string) ([]*usersEntity.Users, error) {
	var results []*usersEntity.Users

	for _, v := range uid {
		var row usersEntity.Users
		r := r.conn.QueryRowContext(ctx, "SELECT * FROM users WHERE account_unique_id=?", v)
		err := r.Scan(
			&row.ID,
			&row.AccountUniqueID,
			&row.Firstname,
			&row.Surename,
			&row.Email,
			&row.Phone,
			&row.CreatedAt,
			&row.UpdatedAt,
			&row.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		if !isDeleted(row.DeletedAt) {
			results = append(results, &row)
		}
	}

	return results, nil
}

func (r *repository) GetAllAcccount(ctx context.Context) ([]*usersEntity.Users, error) {
	var results []*usersEntity.Users

	rows, err := r.conn.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var row usersEntity.Users
		err := rows.Scan(
			&row.ID,
			&row.AccountUniqueID,
			&row.Firstname,
			&row.Surename,
			&row.Email,
			&row.Phone,
			&row.CreatedAt,
			&row.UpdatedAt,
			&row.DeletedAt,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		if !isDeleted(row.DeletedAt) {
			results = append(results, &row)
		}
	}
	defer rows.Close()

	return results, nil
}

func (r *repository) UpdateAccount(ctx context.Context, uid []string, payload []*usersEntity.Users) error {
	for i, v := range payload {
		_, err := r.conn.ExecContext(ctx, "UPDATE users SET firstname=?, surename=?, email=?, phone=?, updated_at=? WHERE account_unique_id=?",
			v.Firstname, v.Surename, v.Email, v.Phone, time.Now().UTC(), uid[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) DeleteAccount(ctx context.Context, uid []string) error {
	for _, v := range uid {
		_, err := r.conn.ExecContext(ctx, "UPDATE users SET deleted_at=? WHERE account_unique_id=?", time.Now().UTC(), v)
		if err != nil {
			return err
		}
	}

	return nil
}
