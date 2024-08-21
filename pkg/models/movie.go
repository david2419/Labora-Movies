package models

type Genre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Movie struct {
	Id               int      `json:"id"`
	Title            string   `json:"original_title"`
	TagLine          string   `json:"Fantasy...beyond your imagination"`
	Overview         string   `json:"overview"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	Genres           []Genre  `json:"genres"`
	Comentarios      []Comentario
}
