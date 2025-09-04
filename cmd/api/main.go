package main

import (
	"RiderApi/internal/handler"
	"RiderApi/internal/repository"
	"RiderApi/internal/service"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	connStr := "user=postgres password='' dbname=KamenRiders host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos")
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("No se pudo verificar la conexión")
	}

	riderRepo := repository.NewRiderRepository(db)
	riderService := service.NewRiderService(riderRepo)
	riderHandler := handler.NewRiderHandler(riderService)

	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Sistema Rider API: !Online!"))
	})

	router.Get("/riders", riderHandler.GetAllRiders)

	fmt.Println("Servidor escuchando el puerto 8080...")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}

}
