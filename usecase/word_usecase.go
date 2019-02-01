package usecase

import (
	"log"

	"github.com/labstack/echo"
	"github.com/wmsuke/WordsOfWisdom/domains/models"
	"github.com/wmsuke/WordsOfWisdom/domains/services"
)

type wordUseCase struct {
}

type WordUseCase interface {
	GetRandomWord(c echo.Context) *models.Word
	GetWord(c echo.Context, wordId int) *models.Word
}

func NewWordUseCase() WordUseCase {
	return &wordUseCase{}
}

func (w *wordUseCase) GetWord(c echo.Context, wordId int) *models.Word {
	cookie, err := c.Cookie("words_userkey")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var v = services.NewWordServices()
	word, err := v.GetWord(wordId, cookie.Value)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return word
}

func (w *wordUseCase) GetRandomWord(c echo.Context) *models.Word {
	cookie, err := c.Cookie("words_userkey")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var v = services.NewWordServices()
	word, err := v.GetRandomWord(cookie.Value)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return word
}
