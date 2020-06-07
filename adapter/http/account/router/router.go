package router

import (
	"github.com/labstack/echo"
	"github.com/wahyurudiyan/authentication/adapter/http/account/handler"
)

type router struct {
	handler handler.Handler
}

type Router interface {
	InitRouter(e *echo.Echo)
}

func NewRouter(handler handler.Handler) Router {
	return &router{handler}
}

func (r *router) InitRouter(e *echo.Echo) {
	routerGroup := e.Group("/api/v1/account")
	routerGroup.POST("/create", r.handler.CreateAccount)
	routerGroup.GET("/get", r.handler.GetAllAccount)
	routerGroup.GET("/get/id", r.handler.GetAccountByID)
	routerGroup.GET("/get/unique", r.handler.GetAccountByUniqueID)
	routerGroup.PUT("/update/:id", r.handler.UpdateAccount)
	routerGroup.DELETE("/delete/:id", r.handler.DeleteAccount)
}
