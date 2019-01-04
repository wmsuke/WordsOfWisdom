package main

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo"
)

func main() {
  e := echo.New()
  e.GET("/words", func(c echo.Context) error {
    return c.String(http.StatusOK, "words, GET!")
  })
  e.Logger.Fatal(e.Start(":1323"))
  host := "localhost"
  port := "1323"
  e.Start(fmt.Sprintf("%v:%v", host, port))
}
