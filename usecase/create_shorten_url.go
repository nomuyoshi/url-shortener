package usecase

import (
	"database/sql"
	"math/big"
	"urlshortener/snowflake"
)

type createShortenURLUseCase struct {
	db *sql.DB
}

func NewCreateShortenURLUseCase(db *sql.DB) *createShortenURLUseCase {
	return &createShortenURLUseCase{db: db}
}

func (uc *createShortenURLUseCase) Execute(longURL string) (string, error) {

	id := snowflake.SnowFlake().Generate()
	shortenURL := big.NewInt(int64(id)).Text(62)

	_, err := uc.db.Exec("INSERT INTO url_mapping (id, long_url, short_url) VALUES (?, ?, ?)", id, longURL, shortenURL)
	if err != nil {
		return "", err
	}

	return shortenURL, nil
}
