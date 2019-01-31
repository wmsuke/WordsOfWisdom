package repositores

import (
	"log"

	"github.com/wmsuke/WordsOfWisdom/domains/models"
)

type wordRepository struct {
}

type WordRepository interface {
	FindOne(ID int) (*models.Words, error)
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

func isFavorite(wordID int, userKey string) (bool, error) {
	sql := `
	SELECT *
	FROM favorites AS f
	INNER JOIN users AS u
	ON f.user_id = u.id
	WHERE f.word_id = ?
	AND u.key = ?
	`
	has, err := engine.SQL(sql, wordID, userKey).Exist()
	if err != nil {
		log.Fatalf("%v", err)
		return false, err
	}
	return has, err
}

func isNice(wordID int, userKey string) (bool, error) {
	sql := `
	SELECT *
	FROM nices AS n
	INNER JOIN users AS u
	ON n.user_id = u.id
	WHERE n.word_id = ?
	AND u.key = ?
	`
	has, err := engine.SQL(sql, wordID, userKey).Exist()
	if err != nil {
		log.Fatalf("%v", err)
		return false, err
	}
	return has, err
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

func (wordRepository *wordRepository) RandomOne(userId int) (*models.Word, error) {
	word := make([]models.Words, 0)
	sql := `
		SELECT * FROM words
		WHERE id=(SELECT (max(id) * random())::int FROM words);
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

	favoriteStatus, err := isFavorite(word[0].Id, "abc1")
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	niceStatus, err := isNice(word[0].Id, "abc1")
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
