package main

import (
	"fmt"
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
	fmt.Println("Se cargaron las configuraciones iniciales de forma exitosa")

	//Database
	//Conexion
	db, err := database.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("No se conectó la BD: %v", err)
	}
	fmt.Println("Se conectó a la base de datos RDS aws de forma exitosa")
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
