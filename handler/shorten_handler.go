package handler

import (
	"math/big"
	"net/http"

	"github.com/labstack/echo/v4"

	"urlshortener/snowflake"
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

	id := snowflake.SnowFlake().Generate()
	shortenURL := big.NewInt(int64(id)).Text(62)
	// TODO DB保存
	res := shortenResponse{
		ShortenURL: shortenURL,
	}

	return c.JSON(http.StatusCreated, res)
}
