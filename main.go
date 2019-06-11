package main

import (
	"fmt"
	"os"

	// "net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wmsuke/WordsOfWisdom/infrastructure/config"
	handlers "github.com/wmsuke/WordsOfWisdom/presenter/handler"
)

func main() {
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	handlers.Router(e)

	// サーバ起動
	if os.Getenv("PORT") == "" {
		var c = config.NewConfig()
		host := c.SEVER.Host
		port := c.SEVER.Port
		e.Start(fmt.Sprintf("%v:%v", host, port))
	} else {
		e.Start(fmt.Sprintf(":%v", os.Getenv("PORT")))
	}
}
