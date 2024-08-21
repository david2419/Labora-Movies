package models

import "database/sql"

type Comentario struct {
	Id         int
	Id_movie   int
	Id_usuario int
	Texto      string
}

func DevuelveComentarios(db *sql.DB, id_movie int) ([]*Comentario, error) {
	query := `
		select * from comentarios where id_movie = ?	
	`
	rows, err := db.Query(query, id_movie)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comentarios []*Comentario

	for rows.Next() {
		var comentario Comentario
		if err := rows.Scan(&comentario.Id, &comentario.Id_movie, &comentario.Id_usuario, &comentario.Texto); err != nil {
			return nil, err
		}
		comentarios = append(comentarios, &comentario)
	}

	return comentarios, nil
}
