package main

import (
	"fmt"
	// "net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wmsuke/backend/handlers"
)

func main() {
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	handlers.Router(e)

	// サーバ起動
	host := "localhost"
	port := "1323"
	e.Start(fmt.Sprintf("%v:%v", host, port))
}
