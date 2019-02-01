package services

import (
	"log"

	"github.com/wmsuke/WordsOfWisdom/domains/repositores"
)

type deleteFavoriteServices struct {
}

type DeleteFavoriteServices interface {
	Delete(wordId int, key string) error
}

func NewDeleteFavoriteServices() DeleteFavoriteServices {
	return &deleteFavoriteServices{}
}

func (a *deleteFavoriteServices) Delete(wordId int, key string) error {
	var v = repositores.NewUserRepository()
	user, err := v.FindOne(key)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}

	var f = repositores.NewFavoriteRepository()
	err = f.Delete(wordId, user.Id)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	return nil
}
