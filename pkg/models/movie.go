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
	Comentarios      []*Comentario
}

func MovieDetails(db *sql.DB, ApiToken string, id int) (*Movie, error) {

	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%v", id)

	req, _ := http.NewRequest("GET", url, nil)

	// header := fmt.Sprintf("Bearer %v", ApiToken)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJiYTliMzUzYzVhMzJhZTA1NTY3YTFmOGEwZTRlMjgwZCIsIm5iZiI6MTcyNDI3MzA2OS45NTg0NDcsInN1YiI6IjY2YzY1MDczYWY4OGMxZjZmMDUyOGZiOSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.LC5Or0sdxSsaEaQvPBZU4c1mLEsvoL5ch9xIl2w46tQ")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error aca:%v", err)
	}
	defer res.Body.Close()

	var movie Movie

	err = json.NewDecoder(res.Body).Decode(&movie)
	if err != nil {
		return nil, fmt.Errorf("error en decode : %v", err)
	}

	if err := incrementaVisualizaciones(db, movie.Id); err != nil {
		return nil, err
	}

	comentarios, err := DevuelveComentarios(db, movie.Id)
	if err != nil {
		return nil, err
	}
	movie.Comentarios = comentarios

	return &movie, nil
}

func incrementaVisualizaciones(db *sql.DB, id int) error {
	query := `
		insert into visualizaciones(id_movie, visualizaciones) 
		values (?, 1)
		on duplicate key update visualizaciones = visualizaciones + 1
	`
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error en incrementoVisualizaciones : %v", err)
	}

	return nil
}
