package models

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}

func GetUserFromCookie(db *sql.DB, r *http.Request, secretKey []byte) (*User, error) {

	//recuperar la cookie con request
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return nil, err
	}

	//decodificar el jwt

	token, err := jwt.ParseWithClaims(cookie.Value, jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
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
