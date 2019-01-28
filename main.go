package main

import (
	"fmt"
	"os"
	// "net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wmsuke/WordsOfWisdom/presenter/handler"
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
		host := "localhost"
		port := "1323"
		e.Start(fmt.Sprintf("%v:%v", host, port))
	} else {
		e.Start(fmt.Sprintf(":%v", os.Getenv("PORT")))
	}
}
