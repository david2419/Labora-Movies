package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Genre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Movie struct {
	Id               int      `json:"id"`
	Title            string   `json:"original_title"`
	TagLine          string   `json:"tagline"`
	Overview         string   `json:"overview"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	Genres           []Genre  `json:"genres"`
}

func MovieDetails(db *sql.DB, ApiToken string, id int) (*Movie, error) {

	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%v", id)

	req, _ := http.NewRequest("GET", url, nil)

	header := fmt.Sprintf("Bearer %v", ApiToken)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", header)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var movie Movie

	err = json.NewDecoder(res.Body).Decode(&movie)
	if err != nil {
		return nil, fmt.Errorf("error en decode : %v", err)
	}

	return &movie, nil
}
