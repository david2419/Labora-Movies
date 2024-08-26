package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"proyectoFinal/pkg/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func RouterHandlers(router *mux.Router, db *sql.DB, apiMovies_access_token string, jwt_secret_key string) {

	//APIS
	router.HandleFunc("/register", Register(db)).Methods("POST")
	router.HandleFunc("/login", Login(db, jwt_secret_key)).Methods("POST")
	router.HandleFunc("/logout", Logout()).Methods("POST")
	router.HandleFunc("/movie/{id}", Movie(db, apiMovies_access_token, jwt_secret_key)).Methods("GET")
	router.HandleFunc("/movie/comentario", AgregarComentario(db, jwt_secret_key)).Methods("POST")

}

// Controladores
func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data map[string]string

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "err.Error()", http.StatusBadRequest)
		}
		log.Println(data)

		//	bcrypt: "go get golang.org/x/crypto/bcrypt", import"golang.org/x/crypto/bcrypt"

		password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
		log.Println(password)

		_, err = db.Exec(
			"INSERT INTO  users (name, email, password) VALUES (?, ?, ?)",
			data["name"], data["email"], password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)
	}
}

func Login(db *sql.DB, jwt_secret_key string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var data map[string]string

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		var user models.User

		err = db.QueryRow("SELECT id, password FROM users WHERE email = ?", data["email"]).Scan(&user.Id, &user.Password)
		if err == sql.ErrNoRows {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//	Verificamos hash Comparar los password
		//https://bcrypt.online/
		err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
		if err != nil {
			http.Error(w, "credenciales incorrectas", http.StatusBadRequest)
			return
		}

		//Generamos un JWT: "go get github.com/dgrijalva/jwt-go"
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(user.Id)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
		})

		token, err := claims.SignedString([]byte(jwt_secret_key))
		if err != nil {
			http.Error(w, "could not login", http.StatusInternalServerError)
			return
		}

		//	Seteo del JWT en un cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true, // agrega capa de seguridad a la cookie para que no se pueda ver desde el lado del cliente
		})

		json.NewEncoder(w).Encode(map[string]string{"message": "success"})

	}
}

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HttpOnly: true,
		})

		json.NewEncoder(w).Encode(map[string]string{"message": "success"})
	}
}

func Movie(db *sql.DB, ApiToken string, jwt_secret_key string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := models.GetUserFromCookie(db, r, jwt_secret_key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//recuperar el movie id
		params := mux.Vars(r)
		idStr := params["id"]

		// Convertir el ID a entero
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		movie, err := models.MovieDetails(db, ApiToken, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// w.WriteHeader().content
		json.NewEncoder(w).Encode(movie)
	}
}

func AgregarComentario(db *sql.DB, jwt_secret_key string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		user, err := models.GetUserFromCookie(db, r, jwt_secret_key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var comentario models.Comentario
		comentario.Id_usuario = int(user.Id)
		if err := json.NewDecoder(r.Body).Decode(&comentario); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := models.CrearComentario(db, comentario.Id_movie, comentario.Id_usuario, comentario.Texto); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(comentario)
	}
}

// func EliminarComentario()http.HandlerFunc{
// 	return func(w http.ResponseWriter, r *http.Request) {

// 	}
// }

// func EditarComentario()http.HandlerFunc{
// 	return func(w http.ResponseWriter, r *http.Request) {

// 	}
// }
