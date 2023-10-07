package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"urlshortener/db"
	"urlshortener/handler"
	"urlshortener/snowflake"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (v *CustomValidator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	db.SetupConnection()

	if err := snowflake.SetSnowFlake(); err != nil {
		panic(err)
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())
	e.POST("/api/shorten_url", handler.CreateShortenURLHandler)
	e.GET("/:short_url", handler.RedirectHandler)
	e.Logger.Fatal(e.Start(":3000"))
}
