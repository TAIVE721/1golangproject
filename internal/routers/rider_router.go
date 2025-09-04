package routers

import (
	"RiderApi/internal/handler"

	"github.com/go-chi/chi/v5"
)

func SetupRiderRoutes(router *chi.Mux, riderHandler *handler.RiderHandler) {

	router.Route("/riders", func(r chi.Router) {
		r.Get("/", riderHandler.GetAllRiders)

	})
}
