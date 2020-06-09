package http

import (
	"github.com/wahyurudiyan/authentication/entity/account"
)

type Request struct {
	ID    []string             `json:"id"`
	Token string               `json:"token"`
	Data  []*account.Users `json:"data"`
}

type Response struct {
	StatusCode int                  `json:"status_code"`
	Message    string               `json:"message"`
	Data       []*account.Users `json:"data"`
}
