package http

import (
	"github.com/wahyurudiyan/authentication/entity/usersEntity"
)

type Request struct {
	ID    []string             `json:"id"`
	Token string               `json:"token"`
	Data  []*usersEntity.Users `json:"data"`
}

type Response struct {
	StatusCode int                  `json:"status_code"`
	Message    string               `json:"message"`
	Data       []*usersEntity.Users `json:"data"`
}
