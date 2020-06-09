package auth

import (
	"time"

	"github.com/wahyurudiyan/authentication/entity/account"

	"github.com/dgrijalva/jwt-go"
)

type Login struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginHistory struct {
	AccountUniqueID string `json:"account_unique_id"`
	Login           `json:"login_info"`
	Role            string    `json:"role"`
	LoginAt         time.Time `json:"login_at`
}

type Register struct {
	*account.Users
}

type AuthClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
