package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"urlshortener/db"
	"urlshortener/usecase"
)

// ShortenHandler 短縮URL生成APIのハンドラー
type shortenRequest struct {
	URL string `json:"url" validate:"http_url"`
}

type shortenResponse struct {
	ShortenURL string `json:"url"`
}

func CreateShortenURLHandler(c echo.Context) error {
	var req shortenRequest
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("failed to bind request. error = %s")
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	uc := usecase.NewCreateShortenURLUseCase(db.DBConn())
	shortenURL, err := uc.Execute(req.URL)
	if err != nil {
		return err
	}

	res := shortenResponse{
		ShortenURL: shortenURL,
	}

	return c.JSON(http.StatusCreated, res)
}
