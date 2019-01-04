package handlers

import (
  "net/http"
  "github.com/labstack/echo"
)

func Router(e *echo.Echo) {
  e.GET("/words", func(c echo.Context) error {
    return c.String(http.StatusOK, "words, GET!")
  })
}
