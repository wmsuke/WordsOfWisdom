package handlers

import (
	"log"
	"net/http"
	"strconv"

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

func validateId(id string) (int, error) {
	return strconv.Atoi(id)
}

func Router(e *echo.Echo) {

	m := e.Group("", UpdateUserKey)
	m.Static("/", "index.html")
	m.Static("/test_form.html", "test_form.html")
	m.Static("/test_complete.html", "test_complete.html")
	e.Static("/static/*", "static")

	c := e.Group("", CheckUserKey)
	c.GET("/words", func(c echo.Context) error {
		var v = usecase.NewWordUseCase()
		return c.JSON(http.StatusOK, v.GetRandomWord(c))
	})
	c.GET("/words/:id", func(c echo.Context) error {
		id, err := validateId(c.Param("id"))
		if err != nil {
			log.Fatal(err)
			return nil
		}
		var v = usecase.NewWordUseCase()
		return c.JSON(http.StatusOK, v.GetWord(c, id))
	})
	c.POST("/words/:id/nice", func(c echo.Context) error {
		id, err := validateId(c.Param("id"))
		if err != nil {
			log.Fatal(err)
			return nil
		}
		var v = usecase.NewNiceUseCase()
		return c.JSON(http.StatusOK, v.Add(c, id))
	})
	c.DELETE("/words/:id/nice", func(c echo.Context) error {
		id, err := validateId(c.Param("id"))
		if err != nil {
			log.Fatal(err)
			return nil
		}
		var v = usecase.NewNiceUseCase()
		return c.JSON(http.StatusOK, v.Delete(c, id))
	})
	c.POST("/words/:id/favorite", func(c echo.Context) error {
		id, err := validateId(c.Param("id"))
		if err != nil {
			log.Fatal(err)
			return nil
		}
		var v = usecase.NewFavoriteUseCase()
		return c.JSON(http.StatusOK, v.Add(c, id))
	})
	c.DELETE("/words/:id/favorite", func(c echo.Context) error {
		id, err := validateId(c.Param("id"))
		if err != nil {
			log.Fatal(err)
			return nil
		}
		var v = usecase.NewFavoriteUseCase()
		return c.JSON(http.StatusOK, v.Delete(c, id))
	})
}
