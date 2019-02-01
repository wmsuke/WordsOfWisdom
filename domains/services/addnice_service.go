package services

import (
	"log"

	"github.com/wmsuke/WordsOfWisdom/domains/repositores"
)

type addNiceServices struct {
}

type AddNiceServices interface {
	Add(wordId int, key string) error
}

func NewAddNiceServices() AddNiceServices {
	return &addNiceServices{}
}

func (a *addNiceServices) Add(wordId int, key string) error {
	var v = repositores.NewUserRepository()
	user, err := v.FindOne(key)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}

	var f = repositores.NewNiceRepository()
	err = f.Insert(wordId, user.Id)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	return nil
}
