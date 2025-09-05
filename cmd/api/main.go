package main

import (
	"RiderApi/internal/handlers"
	"RiderApi/internal/repositories"
	"RiderApi/internal/routers"
	"RiderApi/internal/services"
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	config := "user=postgres password='' dbname=KamenRiders host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("pgx", config)

	if err != nil {
		log.Fatal("Error del servidor")
		return
	}

	if err := db.Ping(); err != nil {
		log.Fatal("No se pudo conectar con la db")
		return
	}

	riderRepo := repositories.NewRiderRepository(db)

	riderService := services.NewRiderService(riderRepo)

	riderHandler := handlers.NewRiderHandler(riderService)

	router := chi.NewRouter()

	routers.SetRiderRouter(router, riderHandler)

	if err := http.ListenAndServe(":8080", router); err != nil {

		log.Fatal("Internal srver error")
		return

	}

}
