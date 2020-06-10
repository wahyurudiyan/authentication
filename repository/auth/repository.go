package auth

import (
	"context"
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"

	accountEntity "github.com/wahyurudiyan/authentication/entity/account"
	authEntity "github.com/wahyurudiyan/authentication/entity/auth"
)

const (
	inputedInvalid  = "inputed username is empty, please fill it first!"
	passwordInvalid = "your password is invalid, plase check again!"
)

type repository struct {
	conn *sql.DB
}

type Repository interface {
	LoginHistory(ctx context.Context, v *authEntity.LoginHistory) error
}

func NewRepository(conn *sql.DB) Repository {
	return &repository{conn}
}

func passwordValidator(pass string, inputed string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(inputed))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) Login(ctx context.Context, v *authEntity.Login) (*accountEntity.Users, error) {
	query := `SELECT * FROM users WHERE username=?`

	if v.Username == "" || v.Password == "" {
		return nil, errors.New(inputedInvalid)
	}

	stmt, err := r.conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, v.Username)
	if err != nil {
		return nil, err
	}

	var result accountEntity.Users
	for rows.Next() {
		err := rows.Scan(
			&result.AccountUniqueID,
			&result.Firstname,
			&result.Surename,
			&result.Username,
			&result.Password,
			&result.Email,
			&result.Phone,
			&result.Role,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.DeletedAt,
		)

		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	valid, _ := passwordValidator(result.Password, v.Password)
	if !valid {
		return nil, errors.New(passwordInvalid)
	}

	return &result, nil
}

func (r *repository) LoginHistory(ctx context.Context, v *authEntity.LoginHistory) error {
	query := `INSERT INTO auth.login_history
	(account_unique_id, username, role, login_at)
	VALUES(?, ?, ?, ?);`

	_, err := r.conn.ExecContext(ctx, query, v.AccountUniqueID, v.Username, v.Role, v.LoginAt)
	if err != nil {
		return err
	}

	return nil
}
