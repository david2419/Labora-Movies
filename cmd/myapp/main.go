package main

import (
	"log"
	"net/http"
	"proyectoFinal/internal/config"
	database "proyectoFinal/internal/dataBase"
	"proyectoFinal/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {

	//Configuraciones

	cfg, err := config.Config()
	if err != nil {
		log.Fatal("No se cargaron las configuraciones iniciales")
	}

	//Database
	//Conexion
	db, err := database.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatal("No se conectó la BD")
	}
	defer db.Close()
	//Creacion de tablas
	if err := database.CreateTablaUsers(db); err != nil {
		log.Fatal(err)
	}
	if err := database.CreateTablaComentarios(db); err != nil {
		log.Fatal(err)
	}
	if err := database.CreateTablaVisualizaciones(db); err != nil {
		log.Fatal(err)
	}

	//Router y handlers
	router := mux.NewRouter()

	//CORS
	// c := cors.New(cors.Options{
	// 	AllowOrigins: "*",
	// 	AllowCredentials: true,
	// })
	//handler := c.Handler(router)

	handlers.RouterHandlers(router, db, cfg.ApiMovies_Access_Token, cfg.JWT_Secret_Key)

	//Listen Server
	log.Printf("Server running on port: %v", cfg.Server_Adress)
	if err := http.ListenAndServe(cfg.Server_Adress, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
