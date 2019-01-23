package main

import (
	"fmt"
	"net/http"
	"os"
	// "net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wmsuke/WordsOfWisdom/presenter/handler"
	"github.com/wmsuke/WordsOfWisdom/usecase"
)

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var v = usecase.NewCheckUseCase()
		err := v.CheckUser(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return next(c)
	}
}

func main() {
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(CustomMiddleware)
	e.Use(middleware.Recover())

	// ルーティング
	handlers.Router(e)

	// サーバ起動
	if os.Getenv("PORT") == "" {
		host := "localhost"
		port := "1323"
		e.Start(fmt.Sprintf("%v:%v", host, port))
	} else {
		e.Start(fmt.Sprintf(":%v", os.Getenv("PORT")))
	}
}
