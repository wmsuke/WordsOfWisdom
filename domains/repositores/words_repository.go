package repositores

import (
	"log"

	"github.com/wmsuke/WordsOfWisdom/domains/models"
)

type wordRepository struct {
}

type WordRepository interface {
	FindOne(wordId int, userId int) (*models.Word, error)
	RandomOne(userId int) (*models.Word, error)
}

func NewWordRepository() WordRepository {
	return &wordRepository{}
}

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

func isFavorite(wordID int, userId int) (bool, error) {
	var favorite = models.Favorites{WordId: wordID, UserId: userId}
	has, err := engine.Exist(&favorite)

	if err != nil {
		log.Fatalf("%v", err)
		return false, err
	}
	return has, err
}

func isNice(wordID int, userId int) (bool, error) {
	var nice = models.Nices{WordId: wordID, UserId: userId}
	has, err := engine.Exist(&nice)
	if err != nil {
		log.Fatalf("%v", err)
		return false, err
	}
	return has, err
}

func (wordRepository *wordRepository) FindOne(wordId int, userId int) (*models.Word, error) {
	var word = models.Words{}
	has, err := engine.Where("id = ?", wordId).Get(&word)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}
	if has {
		niceCount, err := getNiceCount(word.Id)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}

		favoriteCount, err := getFavoriteCount(word.Id)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}

		favoriteStatus, err := isFavorite(word.Id, userId)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}

		niceStatus, err := isNice(word.Id, userId)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}

		return &models.Word{
			ID:              word.Id,
			Word:            word.Word,
			Author:          word.Author,
			Nice:            niceCount,
			Favortite:       favoriteCount,
			NiceStatus:      niceStatus,
			FavortiteStatus: favoriteStatus,
		}, nil

	}

	return nil, nil
}

func (wordRepository *wordRepository) RandomOne(userId int) (*models.Word, error) {
	word := make([]models.Words, 0)
	sql := `
		SELECT * FROM words
		WHERE id=(SELECT (max(id) * random())::int FROM words);
	`

	if err := engine.Sql(sql).Find(&word); err != nil {
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

	favoriteStatus, err := isFavorite(word[0].Id, userId)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	niceStatus, err := isNice(word[0].Id, userId)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	return &models.Word{
		ID:              word[0].Id,
		Word:            word[0].Word,
		Author:          word[0].Author,
		Nice:            niceCount,
		Favortite:       favoriteCount,
		NiceStatus:      niceStatus,
		FavortiteStatus: favoriteStatus,
	}, nil
}
