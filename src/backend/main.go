package main

import (
    "fmt"
    // "net/http"
    "backend/handlers"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
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
