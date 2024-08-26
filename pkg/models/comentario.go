package models

import (
	"database/sql"
	"fmt"
)

type Comentario struct {
	Id         int `json:"id"`
	Id_movie   int `json:"id_movie"`
	Id_usuario int
	Texto      string `json:"comentario"`
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

func CrearComentario(db *sql.DB, id_movie int, id_usuario int, comentario string) error {
	query := `
	INSERT INTO comentarios (id_movie, id_usuario, comentario) 
	VALUES (?, ?, ?);
	`
	_, err := db.Exec(query, id_movie, id_usuario, comentario)
	if err != nil {
		return fmt.Errorf("error al agregar comentario: %v", err)
	}
	return nil
}

func ObtenerComentario(db *sql.DB, idComentario int) (*Comentario, error) {

	var comentario Comentario
	query := `SELECT id, id_movie, id_usuario, comentario FROM comentarios WHERE id = ?`
	err := db.QueryRow(query, idComentario).Scan(&comentario.Id, &comentario.Id_movie, &comentario.Id_usuario, &comentario.Texto)
	if err != nil {
		return nil, err
	}

	// Devuelve el comentario encontrado
	return &comentario, nil
}

func EditarComentario(db *sql.DB, comentarioAeditar *Comentario, Nuevocomentario string) error {
	query := `
		UPDATE comentarios
        SET comentario = ?
        WHERE id_comentario = ?	
	`
	_, err := db.Exec(query, comentarioAeditar.Id, Nuevocomentario)
	if err != nil {
		return fmt.Errorf("error al editar comentario: %v", err)
	}
	return nil
}
