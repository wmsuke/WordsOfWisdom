package services

import (
	"errors"
	"log"

	"github.com/wmsuke/WordsOfWisdom/domains/models"
	"github.com/wmsuke/WordsOfWisdom/domains/repositores"
)

type wordServices struct {
}

type WordServices interface {
	GetWord(wordId int, key string) (*models.Word, error)
	GetRandomWord(key string) (*models.Word, error)
	GetWordSortedFavorite(key string) (*models.Word, error)
	GetWordSortedNice(key string) (*models.Word, error)
}

func NewWordServices() WordServices {
	return &wordServices{}
}

func (w *wordServices) GetWord(wordId int, key string) (*models.Word, error) {
	var u = repositores.NewUserRepository()
	user, err := u.FindOne(key)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	var v = repositores.NewWordRepository()
	word, err := v.FindOne(wordId, user.Id)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("DBエラー")
	}

	return word, nil
}

func (w *wordServices) GetRandomWord(key string) (*models.Word, error) {
	var u = repositores.NewUserRepository()
	user, err := u.FindOne(key)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	var v = repositores.NewWordRepository()
	word, err := v.RandomOne(user.Id)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("DBエラー")
	}

	return word, nil
}

func (w *wordServices) GetWordSortedFavorite(key string) (*models.Word, error) {
	var u = repositores.NewUserRepository()
	user, err := u.FindOne(key)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	var v = repositores.NewWordRepository()
	word, err := v.FondSortedFavoriteOne(user.Id)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("DBエラー")
	}

	return word, nil
}

func (w *wordServices) GetWordSortedNice(key string) (*models.Word, error) {
	var u = repositores.NewUserRepository()
	user, err := u.FindOne(key)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	var v = repositores.NewWordRepository()
	word, err := v.FondSortedNiceOne(user.Id)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("DBエラー")
	}

	return word, nil
}
