package usecase

import (
	"database/sql"
	"errors"
)

type redirectUseCase struct {
	db *sql.DB
}

func NewRedirectUseCase(db *sql.DB) *redirectUseCase {
	return &redirectUseCase{db: db}
}

func (uc *redirectUseCase) Execute(shortURL string) (string, error) {
	var longURL string
	if err := uc.db.QueryRow("SELECT long_url FROM url_mapping WHERE short_url = ?", shortURL).Scan(&longURL); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}

		return "", err
	}

	return longURL, nil
}
