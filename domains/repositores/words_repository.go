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
	FondSortedFavoriteOne(userId int) (*models.Word, error)
	FondSortedNiceOne(userId int) (*models.Word, error)
	Add(word models.Words) (*models.Word, error)
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
	return has, nil
}

func isNice(wordID int, userId int) (bool, error) {
	var nice = models.Nices{WordId: wordID, UserId: userId}
	has, err := engine.Exist(&nice)
	if err != nil {
		log.Fatalf("%v", err)
		return false, err
	}
	return has, nil
}

func (wordRepository *wordRepository) FindOne(wordId int, userId int) (*models.Word, error) {
	var words = models.Words{}
	has, err := engine.Where("id = ?", wordId).Get(&words)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}
	if has {
		word, err := generateWordModels(words, userId)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}
		return word, nil
	}

	return nil, nil
}

func (wordRepository *wordRepository) RandomOne(userId int) (*models.Word, error) {
	words := make([]models.Words, 0)
	sql := `
		SELECT * FROM words
		WHERE id=(SELECT (max(id) * random())::int FROM words);
	`

	if err := engine.Sql(sql).Find(&words); err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	word, err := generateWordModels(words[0], userId)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}
	return word, nil
}

func (wordRepository *wordRepository) FondSortedFavoriteOne(userId int) (*models.Word, error) {
	words := make([]models.Words, 0)
	sql := `
		SELECT * FROM 
		(SELECT word_id, COUNT(word_id) word_count FROM favorites GROUP BY word_id) AS c
		LEFT JOIN words w ON c.word_id = w.id
		ORDER BY word_count DESC;
	`

	if err := engine.Sql(sql).Find(&words); err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	word, err := generateWordModels(words[0], userId)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}
	return word, nil
}

func (wordRepository *wordRepository) FondSortedNiceOne(userId int) (*models.Word, error) {
	words := make([]models.Words, 0)
	sql := `
		SELECT * FROM 
		(SELECT word_id, COUNT(word_id) word_count FROM nices GROUP BY word_id) AS c
		LEFT JOIN words w ON c.word_id = w.id
		ORDER BY word_count DESC;
	`

	if err := engine.Sql(sql).Find(&words); err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	word, err := generateWordModels(words[0], userId)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}
	return word, nil
}

func generateWordModels(word models.Words, userId int) (*models.Word, error) {
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

func (wordRepository *wordRepository) Add(word models.Words) (*models.Word, error) {
	log.Print("step repository 1")
	has, err := engine.InsertOne(word)
	log.Printf("step repository 2 %d", has)
	if err != nil || has < 1 {
		log.Print("step repository error")
		log.Fatalf("%v", err)
		return nil, err
	}
	log.Print("step repository 2")

	// return &word, nil
	return nil, nil
}
