package usecase

import (
	"log"
	"strconv"

	"github.com/labstack/echo"
	"github.com/wmsuke/WordsOfWisdom/domains/models"
	"github.com/wmsuke/WordsOfWisdom/domains/services"
)

type addFavoriteUseCase struct {
}

type AddFavoriteUseCase interface {
	Add(c echo.Context, wordId string) *models.Word
}

func NewAddFavoriteUseCase() AddFavoriteUseCase {
	return &addFavoriteUseCase{}
}

func (u *addFavoriteUseCase) Add(c echo.Context, wordId string) *models.Word {
	cookie, err := c.Cookie("words_userkey")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	var v = services.NewAddFavoriteServices()
	tmpId, _ := strconv.Atoi(wordId)
	v.Add(tmpId, cookie.Value)

	return NewWordUseCase().GetWord(tmpId)
}
