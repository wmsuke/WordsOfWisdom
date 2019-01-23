package usecase

import (
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
}

func NewCheckUseCase() CheckUseCase {
	return &checkUseCase{}
}

func (u *checkUseCase) CheckUser(c echo.Context) error {
	cookie, err := c.Cookie("words_userkey")
	key := random()
	if err != nil {
		cookie = new(http.Cookie)
		cookie.Name = "words_userkey"
		cookie.Value = key
		cookie.Expires = time.Now().Add(744 * time.Hour)
		c.SetCookie(cookie)
	}
	var v = repositores.NewUserRepository()
	v.UpdateUser(key)
	return nil
}

func random() string {
	return uniuri.NewLen(64)
}
