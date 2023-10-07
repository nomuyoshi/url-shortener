package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"urlshortener/snowflake"
)

type shortenURLUseCase struct {
	db *sql.DB
}

func NewShortenURLUseCase(db *sql.DB) *shortenURLUseCase {
	return &shortenURLUseCase{db: db}
}

func (uc *shortenURLUseCase) Execute(longURL string) (string, error) {
	var hash string
	if err := uc.db.
		QueryRow("SELECT short_hash FROM url_mapping WHERE long_url = ?", longURL).
		Scan(&hash); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}

	if hash != "" {
		return newShortURL(hash), nil
	}

	id := snowflake.SnowFlake().Generate()
	hash = big.NewInt(int64(id)).Text(62)

	_, err := uc.db.Exec("INSERT INTO url_mapping (id, long_url, short_hash) VALUES (?, ?, ?)", id, longURL, hash)
	if err != nil {
		return "", err
	}

	return newShortURL(hash), nil
}

func newShortURL(hash string) string {
	return fmt.Sprintf("http://localhost:3000/%s", hash)
}
