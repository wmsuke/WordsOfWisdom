package repositores

import (
	"log"

	"github.com/wmsuke/WordsOfWisdom/domains/models"
)

type wordRepository struct {
}

type WordRepository interface {
	FindOne(ID int) (*models.Words, error)
	RandomOne() (*models.Word, error)
}

func NewWordRepository() WordRepository {
	return &wordRepository{}
}

// func getNiceStatus

func getNiceCount(wordID int) (int64, error) {
	var nice = models.Nices{}
	counts, err := engine.Where("word_id = ?", wordID).Count(&nice)
	if err != nil {
		log.Fatalf("%v", err)
		return 0, err
	}
	return counts, nil
}

func getFavoriteCount(wordID int) (int64, error) {
	var favorite = models.Favorites{}
	counts, err := engine.Where("word_id = ?", wordID).Count(&favorite)
	if err != nil {
		log.Fatalf("%v", err)
		return 0, err
	}
	return counts, nil
}

func (wordRepository *wordRepository) FindOne(id int) (*models.Words, error) {
	var word = models.Words{}
	has, err := engine.Where("id = ?", id).Get(&word)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}
	if has {
		return &word, nil
	}

	return nil, nil
}

func (wordRepository *wordRepository) RandomOne() (*models.Word, error) {
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
		log.Fatalf("%v", err)
		return nil, err
	}

	niceCount, err := getNiceCount(word[0].Id)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	favoriteCount, err := getFavoriteCount(word[0].Id)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	return &models.Word{
		ID:        word[0].Id,
		Word:      word[0].Word,
		Author:    word[0].Author,
		Nice:      niceCount,
		Favortite: favoriteCount,
	}, nil
}
