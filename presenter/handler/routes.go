package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/wmsuke/WordsOfWisdom/usecase"
)

func UpdateUserKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var v = usecase.NewCheckUseCase()
		err := v.CheckUser(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return next(c)
	}
}

func CheckUserKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var v = usecase.NewCheckUseCase()
		err := v.CheckUserKey(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return next(c)
	}
}

func Router(e *echo.Echo) {

	m := e.Group("", UpdateUserKey)
	m.Static("/", "index.html")
	e.Static("/static/*", "static")

	c := e.Group("", CheckUserKey)
	c.GET("/words", func(c echo.Context) error {
		var v = usecase.NewWordUseCase()
		return c.JSON(http.StatusOK, v.GetRandomWord())
	})
	c.GET("/words/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "words id, GET!")
	})
	c.PUT("/words/:id/nice", func(c echo.Context) error {
		return c.String(http.StatusOK, "words id, nice!")
	})
	c.POST("/words/:id/favorite", func(c echo.Context) error {
		return c.String(http.StatusOK, "words id, favorite!")
	})
	c.DELETE("/words/:id/favorite", func(c echo.Context) error {
		return c.String(http.StatusOK, "words id, favorite delete!")
	})
}
