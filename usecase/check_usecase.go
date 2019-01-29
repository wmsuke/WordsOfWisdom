package usecase

import (
	"errors"
	"net/http"
	"time"

	"github.com/dchest/uniuri"
	"github.com/labstack/echo"
	"github.com/wmsuke/WordsOfWisdom/domains/repositores"
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
	cookie, _ := c.Cookie("words_userkey")
	var v = repositores.NewUserRepository()
	has, nil := v.IsOne(cookie.Value)
	if has {
		return nil
	} else {
		return errors.New("キーが存在しない")
	}
}

func (u *checkUseCase) CheckUser(c echo.Context) error {
	cookie, err := c.Cookie("words_userkey")
	key := cookie.Value
	if len(key) == 0 {
		key = random()
		if err != nil {
			cookie = new(http.Cookie)
			cookie.Name = "words_userkey"
			cookie.Value = key
			cookie.Expires = time.Now().Add(744 * time.Hour)
			c.SetCookie(cookie)
		}
	}
	var v = repositores.NewUserRepository()
	v.UpdateUser(key)
	return nil
}

func random() string {
	return uniuri.NewLen(64)
}
