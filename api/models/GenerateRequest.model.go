package models

type GenerateRequest struct {
	OriginalUrl string `json:"originalUrl"`
	ExpiresAfter int `json:"expiresAfter"`
}

