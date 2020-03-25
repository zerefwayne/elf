package models

import (
	"encoding/json"
	"time"
)

type ElfUrl struct {
	originalURL		string
	shortURL		string
	createdAt		time.Time
	expiresAt		time.Time
}

func (elfUrl ElfUrl) Value() string {
	str, _:= json.Marshal(elfUrl)
	return string(str)
}
