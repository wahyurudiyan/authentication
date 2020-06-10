package account

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/wahyurudiyan/authentication/entity/account"
)

type repository struct {
	conn *sql.DB
}
type Repository interface {
	CreateAcccount(ctx context.Context, user []*account.Users) error
	GetAccountByID(ctx context.Context, id []string) ([]*account.Users, error)
	GetAccountByUniqueID(ctx context.Context, uid []string) ([]*account.Users, error)
	GetAllAcccount(ctx context.Context) ([]*account.Users, error)
	UpdateAccount(ctx context.Context, uid []string, payload []*account.Users) error
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

func hashingPassword(pass string) (string, error) {
	password := []byte(pass)

	bPass, err := bcrypt.GenerateFromPassword(password, 16)
	if err != nil {
		return "", err
	}

	return string(bPass), nil
}

func (r *repository) CreateAcccount(ctx context.Context, user []*account.Users) error {
	var arg string
	var args []string

	for _, v := range user {
		role := strings.Join(v.Role, ";")
		hashPassword, _ := hashingPassword(v.Password)

		arg = fmt.Sprintf("('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v')",
			v.AccountUniqueID, v.Firstname, v.Surename, v.Username, hashPassword, v.Email, v.Phone, role, time.Now(), time.Now())
		args = append(args, arg)
	}

	arg = strings.Join(args, ", ")
	query := fmt.Sprintf("INSERT INTO users (account_unique_id, firstname, surename, username, password, email, phone, role, created_at, updated_at) VALUE %s", arg)

	stmt, err := r.conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAccountByID(ctx context.Context, id []string) ([]*account.Users, error) {
	var results []*account.Users

	for _, v := range id {
		var row account.Users
		r := r.conn.QueryRowContext(ctx, "SELECT * FROM users WHERE id=?", v)
		err := r.Scan(
			&row.ID,
			&row.AccountUniqueID,
			&row.Firstname,
			&row.Surename,
			&row.Username,
			&row.Password,
			&row.Email,
			&row.Phone,
			&row.Role,
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

func (r *repository) GetAccountByUniqueID(ctx context.Context, uid []string) ([]*account.Users, error) {
	var results []*account.Users

	for _, v := range uid {
		var row account.Users
		r := r.conn.QueryRowContext(ctx, "SELECT * FROM users WHERE account_unique_id=?", v)
		err := r.Scan(
			&row.ID,
			&row.AccountUniqueID,
			&row.Firstname,
			&row.Surename,
			&row.Username,
			&row.Password,
			&row.Email,
			&row.Phone,
			&row.Role,
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

func (r *repository) GetAllAcccount(ctx context.Context) ([]*account.Users, error) {
	var results []*account.Users

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
		var row account.Users
		err := rows.Scan(
			&row.ID,
			&row.AccountUniqueID,
			&row.Firstname,
			&row.Surename,
			&row.Username,
			&row.Password,
			&row.Email,
			&row.Phone,
			&row.Role,
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

func (r *repository) UpdateAccount(ctx context.Context, uid []string, payload []*account.Users) error {
	for i, v := range payload {
		role := strings.Join(v.Role, ";")
		_, err := r.conn.ExecContext(ctx, "UPDATE users SET firstname=?, surename=?, username=?, password=?, email=?, phone=?, role=?, updated_at=? WHERE account_unique_id=?",
			v.Firstname, v.Surename, v.Username, v.Password, v.Email, v.Phone, role, time.Now().UTC(), uid[i])
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
