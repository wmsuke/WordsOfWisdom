package repositores

import (
	"log"

	"github.com/wmsuke/WordsOfWisdom/domains/models"
)

type favoriteRepository struct {
}

type FavoriteRepository interface {
	Insert(wordId int, userId int) error
}

func NewFavoriteRepository() FavoriteRepository {
	return &favoriteRepository{}
}

func (r *favoriteRepository) Insert(wordId int, userId int) error {
	var favorite = models.Favorites{WordId: wordId, UserId: userId}
	_, err := engine.InsertOne(favorite)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}

	return nil
}
