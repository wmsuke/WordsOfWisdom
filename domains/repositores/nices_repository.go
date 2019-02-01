package repositores

import (
	"log"

	"github.com/wmsuke/WordsOfWisdom/domains/models"
)

type niceRepository struct {
}

type NiceRepository interface {
	Delete(wordId int, userId int) error
	Insert(wordId int, userId int) error
}

func NewNiceRepository() NiceRepository {
	return &niceRepository{}
}

func (r *niceRepository) Insert(wordId int, userId int) error {
	var nice = models.Nices{WordId: wordId, UserId: userId}
	_, err := engine.InsertOne(nice)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}

	return nil
}

func (r *niceRepository) Delete(wordId int, userId int) error {
	var nice = models.Nices{}
	_, err := engine.Where("word_id = ? AND user_id = ?", wordId, userId).Delete(&nice)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}

	return nil
}
