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
	Add(c echo.Context) *models.Word
}

func NewWordUseCase() WordUseCase {
	return &wordUseCase{}
}

func (w *wordUseCase) GetWord(c echo.Context, wordId int) *models.Word {
	var v = services.NewWordServices()
	word, err := v.GetWord(wordId, getCookieValue(c))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return word
}

func (w *wordUseCase) GetRandomWord(c echo.Context) *models.Word {
	var v = services.NewWordServices()
	switch c.QueryParam("sort") {
	case "":
		word, err := v.GetRandomWord(getCookieValue(c))
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return word
	case "star":
		word, err := v.GetWordSortedFavorite(getCookieValue(c))
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return word
	case "nice":
		word, err := v.GetWordSortedNice(getCookieValue(c))
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return word
	default:
		return nil
	}
}

func (w *wordUseCase) Add(c echo.Context) *models.Word {
	var v = services.NewWordServices()
	word, err := v.Add(getCookieValue(c))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return word
}
