package main

import (
    "fmt"
    // "net/http"
    "backend/handlers"
    "github.com/labstack/echo"
)

func main() {
  e := echo.New()
  // e.GET("/words", func(c echo.Context) error {
  //   return c.String(http.StatusOK, "words, GET!")
  // })
  // Router初期化
  handlers.Router(e)

  host := "localhost"
  port := "1323"
  e.Start(fmt.Sprintf("%v:%v", host, port))
}
