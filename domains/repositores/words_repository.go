package repositores

import (
	"log"

	"github.com/wmsuke/WordsOfWisdom/domains/models"
)

type wordRepository struct {
}

type WordRepository interface {
	FindOne(ID int) (*models.Words, error)
}

func NewWordRepository() WordRepository {
	return &wordRepository{}
}

func (wordRepository *wordRepository) FindOne(id int) (*models.Words, error) {
	var word = models.Words{}
	has, err := engine.Where("id = ?", id).Get(&word)
	if err != nil {
		log.Fatalf("データベースの接続に失敗しました。: %v", err)
		return nil, err
	}
	if has {
		return &word, nil
	}

	return nil, nil
}
