package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}

func GetUserFromCookie(db *sql.DB, r *http.Request, secretKey string) (*User, error) {

	//recuperar la cookie con request
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return nil, err
	}

	//decodificar el jwt
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims := token.Claims.(*jwt.StandardClaims)

	//	Consultamos data en la DB mediante el id del User, obtenido del JWT
	var user User

	err = db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", claims.Issuer).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func ModificarUsuario(db *sql.DB, idUsuario int, r *http.Request) error {

	var data map[string]string

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	//verifico qué valores se pasaron para cambiar
	name, okName := data["name"]
	email, okEmail := data["email"]
	pass, okPassword := data["password"]
	log.Println(data)

	if okName {
		query := `UPDATE users
			  SET name = ?
			  where id = ?
			  `

		_, err = db.Exec(query, name, idUsuario)
		if err != nil {
			return fmt.Errorf("error actualizando el nombre del usuario: %v", err)
		}
	}

	if okEmail {
		query := `UPDATE users
			  SET email= ?
			  where id = ?
			  `
		_, err = db.Exec(query, email, idUsuario)
		if err != nil {
			return fmt.Errorf("error actualizando el email del usuario: %v", err)
		}
	}

	if okPassword {
		password, _ := bcrypt.GenerateFromPassword([]byte(pass), 14) //	bcrypt: "go get golang.org/x/crypto/bcrypt", import"golang.org/x/crypto/bcrypt"
		query := `UPDATE users
			  SET password= ?
			  where id = ?
			  `
		_, err = db.Exec(query, password, idUsuario)
		if err != nil {
			return fmt.Errorf("error actualizando el  password del usuario: %v", err)
		}
	}

	return nil
}
