package usecase

import (
	"errors"
	"log"

	"github.com/wmsuke/WordsOfWisdom/domains/models"
	"github.com/wmsuke/WordsOfWisdom/domains/repositores"
)

type wordUseCase struct {
}

type WordUseCase interface {
	GetWord() (*models.Word, error)
}

func NewWordUseCase() WordUseCase {
	return &wordUseCase{}
}

func (w *wordUseCase) GetWord() (*models.Word, error) {
	var v = repositores.NewWordRepository()
	wa, err := v.FindOne(1)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("DBエラー")
	}
	word := &models.Word{
		ID:     wa.Id,
		Word:   wa.Word,
		Author: wa.Author,
	}

	return word, nil
}
