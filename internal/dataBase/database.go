package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(databaseUrl string) (*sql.DB, error) {

	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTablaUsers(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users(
		    id INT AUTO_INCREMENT, 
		    name VARCHAR(100) NOT NULL,
		    email VARCHAR(100) NOT NULL UNIQUE,
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
		FOREIGN KEY(id_usuario) references users(id)
		)
		`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Table users created or already exists")
	return nil
}

func CreateTablaVisualizaciones(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS visualizaciones (
		id INT AUTO_INCREMENT,
		id_movie INT NOT NULL,
		visualizaciones int,
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
