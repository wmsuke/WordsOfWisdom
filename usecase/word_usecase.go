package usecase

import (
	"github.com/wmsuke/WordsOfWisdom/domains/models"
	"github.com/wmsuke/WordsOfWisdom/domains/repositores"
)

type WordUseCase interface {
	GetWord()
}

type wordUseCase struct {
	repositores.WordRepository
}

func (w *wordUseCase) getWord() (*models.Words, error) {
	// ctrl := services.NewCategory()
	// ctrl.Get(1)

	// wa := w.FindOne(1)
	wa, err := w.FindOne(2)
	if err != nil {
		//error log
	}
	return wa, nil

}
func GetWord() *models.Word {
	v := wordUseCase{}
	wa, err := v.getWord()
	if err != nil {
		//error log
	}
	word := &models.Word{
		ID:     wa.ID,
		Word:   wa.Word,
		Author: wa.Author,
	}

	return word

}
