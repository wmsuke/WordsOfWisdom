package services

import (
	"log"

	"github.com/wmsuke/WordsOfWisdom/domains/repositores"
)

type updateUserServices struct {
}

type UpdateUserServices interface {
	Update(key string) error
}

func NewUpdateUserServices() UpdateUserServices {
	return &updateUserServices{}
}

func (a *updateUserServices) Update(key string) error {
	var v = repositores.NewUserRepository()
	has, err := v.IsOne(key)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}

	if has {
		err = v.UpdateUser(key)
		if err != nil {
			log.Fatalf("%v", err)
			return err
		}
	} else {
		err = v.InsertUser(key)
		if err != nil {
			log.Fatalf("%v", err)
			return err
		}
	}

	return nil
}
