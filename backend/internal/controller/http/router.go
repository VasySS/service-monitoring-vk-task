package http

import (
	"net/http"

	"github.com/VasySS/service-monitoring-vk-task/backend/internal/controller/http/v1/statuses"
	"github.com/go-chi/chi/v5"
)

func NewRouter(statusesService statuses.Service) *chi.Mux {
	r := chi.NewRouter()

	sh := statuses.NewHandler(statusesService)

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/statuses", sh.Routes())
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return r
}
