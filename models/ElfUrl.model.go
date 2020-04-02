package models

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"time"
)

// Helper Methods

// REQUIRED: Requires Generation of normalized URL

func GenerateShortUrl(originalUrl string) string {

	md5Sum := md5.Sum([]byte(originalUrl))
	md5Hash := hex.EncodeToString(md5Sum[:])
	shortUrl := md5Hash[:6]

	return shortUrl
}

type ElfUrl struct {
	OriginalURL		string		`json:"originalUrl" db:"original_url"`
	ShortURL		string		`json:"shortUrl" db:"short_url"`
	CreatedAt		time.Time	`json:"createdAt" db:"created_at"`
	ExpiresAt		time.Time	`json:"expiresAfter" db:"expires_at"`
	HasExpired		bool		`json:"hasExpired" db:"has_expired"`
}

func (newElfUrl ElfUrl) Value() string {
	str, _:= json.Marshal(newElfUrl)
	return string(str)
}

func (newElfUrl *ElfUrl) ParseForm(requestBody *GenerateRequest) {

	newElfUrl.OriginalURL = requestBody.OriginalUrl
	newElfUrl.CreatedAt = time.Now()
	newElfUrl.ExpiresAt = time.Now().Add(time.Second * time.Duration(requestBody.ExpiresAfter))
	newElfUrl.HasExpired = false

}