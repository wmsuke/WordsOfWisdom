package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/wmsuke/WordsOfWisdom/domains/models"
)

func Router(e *echo.Echo) {
	e.Static("/", "index.html")
	e.Static("/static/*", "static")

	e.GET("/words", func(c echo.Context) error {
		// engine, _ := xorm.NewEngine("mysql", "root:admin@/tcp(127.0.0.1:3306)/WordsOfWisdom")
		// words := make([]models.Words, 0)
		// engine.Find(&words)
		// return c.String(http.StatusOK, "words, GET!")
		var words models.Words
		words.Word = "人生において、人から与えられたものは何もなかった。いつも自分で取りにいかなければならない。"
		words.Author = "ケビン・デュラント"
		return c.JSON(http.StatusOK, words)
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