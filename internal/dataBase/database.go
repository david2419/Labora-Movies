package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(databaseUrl string) (*sql.DB, error) {

	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("error abriendo la base de datos: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error haciendo ping con la base de datos: %v", err)
	}
	return db, nil
}

func CreateTablaUsers(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users(
		    id INT AUTO_INCREMENT, 
		    name VARCHAR(100) NOT NULL,
		    email VARCHAR(100) NOT NULL UNIQUE,
			apiToken BLOB NOT NULL,
			password BLOB NOT NULL,
		    PRIMARY KEY(id)
		)
       `

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Table users created or already exists")
	return nil
}

func CreateTablaComentarios(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS comentarios (
		id INT AUTO_INCREMENT,
		id_movie INT NOT NULL,
		id_usuario INT NOT NULL,
		comentario VARCHAR(1000) NOT NULL,
		PRIMARY KEY(id),
		FOREIGN KEY(id_usuario) references users(id),
		FOREIGN KEY(id_movie) references visualizaciones(id_movie)
		)
		`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Table comentarios created or already exists")
	return nil
}

func CreateTablaVisualizaciones(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS visualizaciones (
		id_movie INT NOT NULL UNIQUE,
		visualizaciones int,
		PRIMARY KEY(id_movie)
		)	
		`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Table visualizaciones created or already exists")
	return nil
}
