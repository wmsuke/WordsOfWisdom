package services

import (
	"log"

	"github.com/wmsuke/WordsOfWisdom/domains/repositores"
)

type deleteNiceServices struct {
}

type DeleteNiceServices interface {
	Delete(wordId int, key string) error
}

func NewDeleteNiceServices() DeleteNiceServices {
	return &deleteNiceServices{}
}

func (a *deleteNiceServices) Delete(wordId int, key string) error {
	var v = repositores.NewUserRepository()
	user, err := v.FindOne(key)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}

	var f = repositores.NewNiceRepository()
	err = f.Delete(wordId, user.Id)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	return nil
}
