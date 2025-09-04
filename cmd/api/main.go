package main

import (
	"RiderApi/internal/handler"
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

	riderHandler := handler.NewRiderHandler(db)

	router := chi.NewRouter()

	router.Get("/riders", riderHandler.GetAllRiders)

	fmt.Println("Servidor escuchando el puerto 8080...")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}

}
