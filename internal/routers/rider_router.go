package routers

import (
	"RiderApi/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func SetRiderRouter(router *chi.Mux, handler handlers.RiderHandler) {
	router.Route("/riders", func(r chi.Router) {
		r.Get("/", handler.GetAll)
		r.Post("/", handler.Post)
		r.Patch("/{Rider_id}", handler.Patch)
		r.Delete("/{Rider_id}", handler.Delete)
	})
}
