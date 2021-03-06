package usecase

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dchest/uniuri"
	"github.com/labstack/echo"
	"github.com/wmsuke/WordsOfWisdom/domains/repositores"
	"github.com/wmsuke/WordsOfWisdom/domains/services"
)

type checkUseCase struct {
}

type CheckUseCase interface {
	CheckUser(c echo.Context) error
	CheckUserKey(c echo.Context) error
}

func NewCheckUseCase() CheckUseCase {
	return &checkUseCase{}
}

func (u *checkUseCase) CheckUserKey(c echo.Context) error {
	var v = repositores.NewUserRepository()
	has, nil := v.IsOne(getCookieValue(c))
	if has {
		return nil
	} else {
		return errors.New("キーが存在しない")
	}
}

func (u *checkUseCase) CheckUser(c echo.Context) error {
	cookie, err := c.Cookie("words_userkey")
	key := ""
	if err != nil {
		key = random()
		if err != nil {
			cookie = new(http.Cookie)
			cookie.Name = "words_userkey"
			cookie.Value = key
			cookie.Expires = time.Now().Add(744 * time.Hour)
			c.SetCookie(cookie)
		}
	} else {
		key = cookie.Value
	}
	var v = services.NewUpdateUserServices()
	err = v.Update(key)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func random() string {
	return uniuri.NewLen(64)
}
