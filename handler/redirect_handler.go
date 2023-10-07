package handler

import (
	"net/http"
	"urlshortener/db"
	"urlshortener/usecase"

	"github.com/labstack/echo/v4"
)

type redirectRequest struct {
	Hash string `param:"hash" validate:"required"`
}

func RedirectHandler(c echo.Context) error {
	var req redirectRequest
	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("failed to bind request, error: %s", err.Error())
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	url, err := usecase.NewRedirectUseCase(db.DBConn()).Execute(req.Hash)
	// 500 や 404 ページ作ってリダイレクトした方が良い
	if err != nil {
		c.Logger().Errorf("internal server error: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "予期せぬエラーが発生しました")
	}

	if url == "" {
		return echo.NewHTTPError(http.StatusNotFound, "存在しないURLです")
	}

	return c.Redirect(http.StatusPermanentRedirect, url)
}
