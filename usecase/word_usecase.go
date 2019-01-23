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
	GetRandomWord() *models.Word
	GetWord(id int) *models.Word
}

func NewWordUseCase() WordUseCase {
	return &wordUseCase{}
}

func getWord(id int) (*models.Word, error) {
	var v = repositores.NewWordRepository()
	wa, err := v.FindOne(id)
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

func getRandomWord() (*models.Word, error) {
	var v = repositores.NewWordRepository()
	wa, err := v.RandomOne()
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

func (w *wordUseCase) GetWord(id int) *models.Word {
	word, err := getWord(id)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return word
}

func (w *wordUseCase) GetRandomWord() *models.Word {
	word, err := getRandomWord()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return word
}
