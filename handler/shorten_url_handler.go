package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"urlshortener/db"
	"urlshortener/usecase"
)

// ShortenHandler 短縮URL生成APIのハンドラー
type shortenURLRequest struct {
	URL string `json:"url" validate:"http_url"`
}

type shortenURLResponse struct {
	ShortURL string `json:"url"`
}

func ShortenURLHandler(c echo.Context) error {
	var req shortenURLRequest
	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("failed to bind request, error = %s", err.Error())
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	uc := usecase.NewShortenURLUseCase(db.DBConn())
	shortURL, err := uc.Execute(req.URL)
	if err != nil {
		c.Logger().Errorf("internal server error, error = %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "予期せぬエラーが発生しました")
	}

	res := shortenURLResponse{
		ShortURL: shortURL,
	}

	return c.JSON(http.StatusCreated, res)
}
