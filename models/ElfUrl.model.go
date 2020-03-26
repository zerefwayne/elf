package models

import (
	"encoding/json"
	"time"
)

type ElfUrl struct {
	OriginalURL		string
	ShortURL		string
	CreatedAt		time.Time
	ExpiresAt		time.Time
}

func (elfUrl ElfUrl) Value() string {
	str, _:= json.Marshal(elfUrl)
	return string(str)
}
