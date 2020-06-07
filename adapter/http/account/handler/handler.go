package handler

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	server "github.com/wahyurudiyan/authentication/adapter/http"
	"github.com/wahyurudiyan/authentication/usecase/account"
)

const (
	paramError    = "parameter(s) is empty, please check your url parameter(s)"
	bindingError  = "unable to bind your payload, please check it first"
	idsEmptyError = "no more id, please check your input id"
)

type handler struct {
	usecase account.Usecase
}

type Handler interface {
	CreateAccount(e echo.Context) error
	GetAccountByID(e echo.Context) error
	GetAccountByUniqueID(e echo.Context) error
	GetAllAccount(e echo.Context) error
	UpdateAccount(e echo.Context) error
	DeleteAccount(e echo.Context) error
}

func NewAccountHandler(usecase account.Usecase) Handler {
	return &handler{usecase}
}

func (h *handler) CreateAccount(e echo.Context) error {
	var err error
	var request server.Request
	ctx := context.Background()

	if err = e.Bind(&request); err != nil {
		log.Printf("[ERR] context: bind, handler: CreateAccount, msg: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, bindingError)
	}

	if err = h.usecase.CreateAcccount(ctx, request.Data); err != nil {
		log.Printf("[ERR] context: create users account, handler: CreateAccount, msg: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := server.Response{
		StatusCode: http.StatusOK,
		Message:    "OK",
	}

	return e.JSON(http.StatusOK, response)
}

func (h *handler) GetAccountByID(e echo.Context) error {
	var err error
	var request server.Request
	ctx := context.Background()

	if err = e.Bind(&request); err != nil {
		log.Printf("[ERR] context: bind, handler: CreateAccount, msg: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, bindingError)
	}

	users, err := h.usecase.GetAccountByID(ctx, request.ID)
	if err != nil {
		log.Printf("[ERR] context: get users account, handler: GetAccountByID, msg: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	response := server.Response{
		StatusCode: http.StatusOK,
		Message:    "OK",
		Data:       users,
	}

	return e.JSON(http.StatusOK, response)
}

func (h *handler) GetAccountByUniqueID(e echo.Context) error {
	var err error
	var request server.Request
	ctx := context.Background()

	if err = e.Bind(&request); err != nil {
		log.Printf("[ERR] context: bind, handler: CreateAccount, msg: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, bindingError)
	}

	users, err := h.usecase.GetAccountByUniqueID(ctx, request.ID)
	if err != nil {
		log.Printf("[ERR] context: get users account, handler: GetAccountByUniqueID, msg: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := server.Response{
		StatusCode: http.StatusOK,
		Message:    "OK",
		Data:       users,
	}

	return e.JSON(http.StatusOK, response)
}

func (h *handler) GetAllAccount(e echo.Context) error {
	var err error
	ctx := context.Background()

	users, err := h.usecase.GetAllAcccount(ctx)
	if err != nil {
		log.Printf("[ERR] context: get all users account, handler: GetAllAccount, msg: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := server.Response{
		StatusCode: http.StatusOK,
		Message:    "OK",
		Data:       users,
	}

	return e.JSON(http.StatusOK, response)
}

func (h *handler) UpdateAccount(e echo.Context) error {
	var err error
	var request server.Request
	ctx := context.Background()

	if err = e.Bind(&request); err != nil {
		log.Printf("[ERR] context: bind, handler: UpdateAccount, msg: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, bindingError)
	}

	param := e.Param("id")
	if param == "" {
		log.Printf("[ERR] context: get param, handler: UpdateAccount, msg: %s", paramError)
		return echo.NewHTTPError(http.StatusBadRequest, paramError)
	}

	id := strings.Split(param, ";")
	err = h.usecase.UpdateAccount(ctx, id, request.Data)
	if err != nil {
		log.Printf("[ERR] context: updating, handler: UpdateAccount, msg: %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := server.Response{
		StatusCode: http.StatusOK,
		Message:    "OK",
	}

	return e.JSON(http.StatusOK, response)
}

func (h *handler) DeleteAccount(e echo.Context) error {
	var err error
	var request server.Request
	ctx := context.Background()

	if err = e.Bind(&request); err != nil {
		log.Printf("[ERR] context: bind, handler: DeleteAccount, msg: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, bindingError)
	}

	param := e.Param("id")
	if param == "" {
		log.Printf("[ERR] context: get param, handler: DeleteAccount, msg: %s", paramError)
		return echo.NewHTTPError(http.StatusBadRequest, paramError)
	}

	id := strings.Split(param, ";")
	err = h.usecase.DeleteAccount(ctx, id)
	if err != nil {
		log.Printf("[ERR] context: deleting, handler: DeleteAccount, msg: %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := server.Response{
		StatusCode: http.StatusOK,
		Message:    "OK",
	}

	return e.JSON(http.StatusOK, response)
}
