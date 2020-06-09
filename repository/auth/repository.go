package auth

import (
	"context"
	"database/sql"

	"github.com/wahyurudiyan/users-auth/entity/auth"
)

type repository struct {
	conn *sql.DB
}

type Repository interface {
	LoginHistory(ctx context.Context, v *auth.LoginHistory) error
}

func NewRepository(conn *sql.DB) Repository {
	return &repository{conn}
}

func (r *repository) LoginHistory(ctx context.Context, v *auth.LoginHistory) error {
	query := `INSERT INTO auth.login_history
	(account_unique_id, username, role, login_at)
	VALUES(?, ?, ?, ?);`

	_, err := r.conn.ExecContext(ctx, query, v.AccountUniqueID, v.Username, v.Role, v.LoginAt)
	if err != nil {
		return err
	}

	return nil
}
