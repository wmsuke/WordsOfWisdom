package main

import (
    "fmt"
    // "net/http"
    "backend/handlers"
    "github.com/labstack/echo"
)

func main() {
  e := echo.New()
  handlers.Router(e)

  host := "localhost"
  port := "1323"
  e.Start(fmt.Sprintf("%v:%v", host, port))
}
