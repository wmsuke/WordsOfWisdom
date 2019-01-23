package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/wmsuke/WordsOfWisdom/usecase"
)

func Router(e *echo.Echo) {
	e.Static("/", "index.html")
	e.Static("/static/*", "static")

	e.GET("/words", func(c echo.Context) error {
		var v = usecase.NewWordUseCase()
		return c.JSON(http.StatusOK, v.GetRandomWord())
	})
	e.GET("/words/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "words id, GET!")
	})
	e.PUT("/words/:id/nice", func(c echo.Context) error {
		return c.String(http.StatusOK, "words id, nice!")
	})
	e.POST("/words/:id/favorite", func(c echo.Context) error {
		return c.String(http.StatusOK, "words id, favorite!")
	})
	e.DELETE("/words/:id/favorite", func(c echo.Context) error {
		return c.String(http.StatusOK, "words id, favorite delete!")
	})
}
