package usecase

import (
	"github.com/labstack/echo"
	"github.com/wmsuke/WordsOfWisdom/domains/models"
	"github.com/wmsuke/WordsOfWisdom/domains/services"
)

type niceUseCase struct {
}

type NiceUseCase interface {
	Add(c echo.Context, wordId int) *models.Word
	Delete(c echo.Context, wordId int) *models.Word
}

func NewNiceUseCase() NiceUseCase {
	return &niceUseCase{}
}

func (u *niceUseCase) Add(c echo.Context, wordId int) *models.Word {
	var v = services.NewAddNiceServices()
	v.Add(wordId, getCookieValue(c))
	return NewWordUseCase().GetWord(c, wordId)
}

func (u *niceUseCase) Delete(c echo.Context, wordId int) *models.Word {
	var v = services.NewDeleteNiceServices()
	v.Delete(wordId, getCookieValue(c))
	return NewWordUseCase().GetWord(c, wordId)
}
