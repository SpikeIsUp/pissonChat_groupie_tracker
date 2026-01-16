package meme

import (
	"encoding/json"
	"net/http"
)

type Meme struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetMemes() ([]Meme, error) {
	resp, err := http.Get("https://api.memegen.link/templates")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var memes []Meme
	err = json.NewDecoder(resp.Body).Decode(&memes)
	return memes, err
}
