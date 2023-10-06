package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

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
	if err := snowflake.SetSnowFlake(); err != nil {
		panic(err)
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/api/shorten_url", handler.CreateShortenURLHandler)
	e.Logger.Fatal(e.Start(":3000"))
}
