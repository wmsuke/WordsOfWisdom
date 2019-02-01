package usecase

import (
	"log"
	"strconv"

	"github.com/labstack/echo"
	"github.com/wmsuke/WordsOfWisdom/domains/models"
	"github.com/wmsuke/WordsOfWisdom/domains/services"
)

type favoriteUseCase struct {
}

type FavoriteUseCase interface {
	Add(c echo.Context, wordId string) *models.Word
	Delete(c echo.Context, wordId string) *models.Word
}

func NewFavoriteUseCase() FavoriteUseCase {
	return &favoriteUseCase{}
}

func (u *favoriteUseCase) Add(c echo.Context, wordId string) *models.Word {
	cookie, err := c.Cookie("words_userkey")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	var v = services.NewAddFavoriteServices()
	tmpId, _ := strconv.Atoi(wordId)
	v.Add(tmpId, cookie.Value)

	return NewWordUseCase().GetWord(c, tmpId)
}

func (u *favoriteUseCase) Delete(c echo.Context, wordId string) *models.Word {
	cookie, err := c.Cookie("words_userkey")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	var v = services.NewDeleteFavoriteServices()
	tmpId, _ := strconv.Atoi(wordId)
	v.Delete(tmpId, cookie.Value)

	return NewWordUseCase().GetWord(c, tmpId)
}
