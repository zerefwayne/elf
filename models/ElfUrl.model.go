package models

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// Helper Methods

// REQUIRED: Requires Generation of normalized URL

func generateShortUrl(originalUrl string) string {

	md5Sum := md5.Sum([]byte(originalUrl))
	md5Hash := hex.EncodeToString(md5Sum[:])
	shortUrl := md5Hash[:6]

	return shortUrl
}

type ElfUrl struct {
	OriginalURL		string		`json:"original" db:"original_url"`
	ShortURL		string		`json:"short" db:"short_url"`
	CreatedAt		time.Time	`json:"created_at" db:"created_at"`
	ExpiresAt		time.Time	`json:"expires_at" db:"expires_at"`
	HasExpired		bool		`json:"has_expired" db:"has_expired"`
}

func (newElfUrl ElfUrl) Value() string {
	str, _:= json.Marshal(newElfUrl)
	return string(str)
}

func (newElfUrl *ElfUrl) ParseForm(r *http.Request) {

	newElfUrl.OriginalURL = r.FormValue("originalUrl")
	newElfUrl.CreatedAt = time.Now()

	newElfUrl.ShortURL = generateShortUrl(newElfUrl.OriginalURL)

	stringToDuration, _ := strconv.ParseInt(r.FormValue("expiresAfter"), 10, 64)
	newElfUrl.ExpiresAt = time.Now().Add(time.Second * time.Duration(stringToDuration))
	newElfUrl.HasExpired = false

}