package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/wmsuke/WordsOfWisdom/usecase"
)

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var v = usecase.NewCheckUseCase()
		err := v.CheckUser(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return next(c)
	}
}

func Router(e *echo.Echo) {

	m := e.Group("", CustomMiddleware)
	m.Static("/", "index.html")
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
