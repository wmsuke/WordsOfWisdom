package repositores

import (
	"log"

	"github.com/wmsuke/WordsOfWisdom/domains/models"
)

type wordRepository struct {
}

type WordRepository interface {
	FindOne(ID int) (*models.Words, error)
	RandomOne() (*models.Words, error)
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

func (wordRepository *wordRepository) RandomOne() (*models.Words, error) {
	word := make([]models.Words, 0)
	sql := `
	SELECT s.*
	FROM words AS s
	INNER JOIN (
	  SELECT CEIL(RAND() * (SELECT MAX(id) FROM words)) AS id
	) AS tmp ON s.id >= tmp.id
	ORDER BY s.id
	LIMIT 1
	`
	err := engine.Sql(sql).Find(&word)
	if err != nil {
		log.Fatalf("データベースの接続に失敗しました。: %v", err)
		return nil, err
	}

	return &models.Words{
		Id:     word[0].Id,
		Word:   word[0].Word,
		Author: word[0].Author,
	}, nil
}
