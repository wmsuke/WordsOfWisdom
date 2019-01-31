package usecase

import (
	"log"

	"github.com/labstack/echo"
	"github.com/wmsuke/WordsOfWisdom/domains/services"
)

type addFavoriteUseCase struct {
}

type AddFavoriteUseCase interface {
	Add(c echo.Context, wordId string) error
}

func NewAddFavoriteUseCase() AddFavoriteUseCase {
	return &addFavoriteUseCase{}
}

func (u *addFavoriteUseCase) Add(c echo.Context, wordId string) error {
	cookie, err := c.Cookie("words_userkey")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	var v = services.NewAddFavoriteServices()
	return v.Add(wordId, cookie.Value)
}
