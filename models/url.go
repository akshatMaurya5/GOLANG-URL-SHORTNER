package models

type URL struct {
	ID       int    `json:"id"`
	LongURL  string `json:"url"`
	ShortURL string `json:"shorturl"`
	HitCount int    `json:"hit_count"`
}
