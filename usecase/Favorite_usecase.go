package usecase

import (
	"github.com/labstack/echo"
	"github.com/wmsuke/WordsOfWisdom/domains/models"
	"github.com/wmsuke/WordsOfWisdom/domains/services"
)

type favoriteUseCase struct {
}

type FavoriteUseCase interface {
	Add(c echo.Context, wordId int) *models.Word
	Delete(c echo.Context, wordId int) *models.Word
}

func NewFavoriteUseCase() FavoriteUseCase {
	return &favoriteUseCase{}
}

func (u *favoriteUseCase) Add(c echo.Context, wordId int) *models.Word {
	var v = services.NewAddFavoriteServices()
	v.Add(wordId, getCookieValue(c))
	return NewWordUseCase().GetWord(c, wordId)
}

func (u *favoriteUseCase) Delete(c echo.Context, wordId int) *models.Word {
	var v = services.NewDeleteFavoriteServices()
	v.Delete(wordId, getCookieValue(c))
	return NewWordUseCase().GetWord(c, wordId)
}

func getCookieValue(c echo.Context) string {
	cookie, _ := c.Cookie("words_userkey")
	return cookie.Value
}
