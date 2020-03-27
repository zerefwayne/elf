package database

import (
	"context"
	"errors"
	"github.com/zerefwayne/elf/config"
	"github.com/zerefwayne/elf/models"
	"log"
)

func InsertURL(url *models.ElfUrl) error {

	query := "INSERT INTO shorturl(original_url, short_url, expires_at, created_at, has_expired) VALUES($1, $2, $3, $4, $5)"

	_, insertError := config.DB.Exec(context.Background(), query, url.OriginalURL, url.ShortURL, url.ExpiresAt, url.CreatedAt, url.HasExpired)

	return insertError

}

func FetchOriginalURL(shortUrl string) (string, error) {

	query := `SELECT original_url FROM shorturl WHERE short_url = '`+shortUrl+`' LIMIT 1;`

	//log.Println(query)
	log.Printf("Redirect query for %s", shortUrl)

	rows, queryError := config.DB.Query(context.Background(), query)

	if queryError != nil {
		return "", queryError
	} else {

		defer rows.Close()

		for rows.Next() {
			originalUrl := ""
			err := rows.Scan(&originalUrl)

			if err != nil {
				return "", err
			} else {
				return originalUrl, nil
			}

		}
		return "", errors.New("no Url Found")
	}
}

