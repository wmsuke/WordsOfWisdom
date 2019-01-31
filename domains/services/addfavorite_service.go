package services

import (
	"log"

	"github.com/wmsuke/WordsOfWisdom/domains/repositores"
)

type addFavoriteServices struct {
}

type AddFavoriteServices interface {
	Add(wordId int, key string) error
}

func NewAddFavoriteServices() AddFavoriteServices {
	return &addFavoriteServices{}
}

func (a *addFavoriteServices) Add(wordId int, key string) error {
	var v = repositores.NewUserRepository()
	user, err := v.FindOne(key)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}

	var f = repositores.NewFavoriteRepository()
	err = f.Insert(wordId, user.Id)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	return nil
}
