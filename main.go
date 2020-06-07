package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/wahyurudiyan/authentication/adapter/http/account/handler"
	"github.com/wahyurudiyan/authentication/adapter/http/account/router"
	"github.com/wahyurudiyan/authentication/adapter/mysql"
	"github.com/wahyurudiyan/authentication/config"
	repository "github.com/wahyurudiyan/authentication/repository/account"
	usecase "github.com/wahyurudiyan/authentication/usecase/account"
)

func main() {
	e := echo.New()
	cfg := config.GetConfiguration()

	sqlConnection, err := mysql.NewConnection(cfg)
	if err != nil {
		log.Printf("[ERR] context: main.sqlConnection, msg: %s", err.Error())
	}

	repo := repository.NewAccountRepository(sqlConnection)
	ucase := usecase.NewAccountUsecase(repo)
	handler := handler.NewAccountHandler(ucase)
	router := router.NewRouter(handler)

	router.InitRouter(e)
	log.Fatal(e.Start(fmt.Sprintf(":%s", cfg.PortHTTP)))
}
