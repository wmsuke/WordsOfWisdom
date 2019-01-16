package usecase

import (
	"errors"
	"log"

	"github.com/wmsuke/WordsOfWisdom/domains/models"
	"github.com/wmsuke/WordsOfWisdom/domains/repositores"
)

type WordUseCase interface {
	GetWord() *models.Word
}

type wordUseCase struct {
	WordRepository repositores.WordRepository
}

func (w *wordUseCase) getWord() (*models.Words, error) {
	// ctrl := services.NewCategory()
	// ctrl.Get(1)

	// wa := w.FindOne(1)
	// wa, err := w.FindOne(2)
	// if err != nil {
	// 	return nil, errors.New("DBエラー")
	// }
	// return wa, nil
	was, err := w.WordRepository.FindOne(1)
	if was == nil {

	}
	if err != nil {
		return nil, errors.New("DBエラー")
	}

	wa := &models.Words{
		Id:     1,
		Word:   "test",
		Author: "mark",
	}
	return wa, nil
}

// func GetWord() *models.Word {
func (w *wordUseCase) GetWord() (*models.Word, error) {
	// v := wordUseCase{}
	// wa, err := v.getWord()
	wa, err := w.WordRepository.FindOne(1)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("DBエラー")
	}
	word := &models.Word{
		ID:     wa.Id,
		Word:   wa.Word,
		Author: wa.Author,
	}
	// word := &models.Word{
	// 	ID:     1,
	// 	Word:   "test",
	// 	Author: "mark",
	// }

	return word, nil

}
