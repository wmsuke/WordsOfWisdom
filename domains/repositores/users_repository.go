package repositores

import (
	"log"
	"time"

	"github.com/wmsuke/WordsOfWisdom/domains/models"
)

const layout = "2006/01/02 15:04:05"

type userRepository struct {
}

type UserRepository interface {
	UpdateUser(Key string) error
	IsOne(Key string) (bool, error)
	FindOne(Key string) (*models.Users, error)
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u *userRepository) IsOne(Key string) (bool, error) {
	var user = models.Users{}
	has, err := engine.Where("`key` = ?", Key).Exist(&user)
	if err != nil {
		log.Fatalf("%v", err)
		return false, err
	}
	return has, nil
}

func (u *userRepository) FindOne(Key string) (*models.Users, error) {
	var user = models.Users{}
	_, err := engine.Where("`key` = ?", Key).Get(&user)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) UpdateUser(Key string) error {
	var user = models.Users{}
	_, err := engine.Where("`key` = ?", Key).Delete(&user)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	user.Key = Key
	user.AccessDate = time.Now()
	_, err = engine.Insert(&user)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	return nil
}
