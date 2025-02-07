package statuses

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/VasySS/service-monitoring-vk-task/backend/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Service interface {
	NewContainerStatuses(ctx context.Context, statuses []entity.ContainerStatus) error
	ContainerStatuses(ctx context.Context) ([]entity.ContainerStatus, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h Handler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.newContainerStatuses)
	r.Get("/", h.getContainerStatuses)

	return r
}

func (h Handler) newContainerStatuses(w http.ResponseWriter, r *http.Request) {
	var statuses []entity.ContainerStatus

	if err := render.DecodeJSON(r.Body, &statuses); err != nil {
		http.Error(w, "error decoding statuses body", http.StatusBadRequest)
		slog.Warn("error decoding statuses body", slog.Any("error", err))

		return
	}

	if err := h.service.NewContainerStatuses(r.Context(), statuses); err != nil {
		http.Error(w, "error saving statuses", http.StatusInternalServerError)
		slog.Warn("error saving statuses", slog.Any("error", err))

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) getContainerStatuses(w http.ResponseWriter, r *http.Request) {
	statuses, err := h.service.ContainerStatuses(r.Context())
	if err != nil {
		http.Error(w, "error getting statuses", http.StatusInternalServerError)
		slog.Warn("error getting statuses", slog.Any("error", err))

		return
	}

	render.JSON(w, r, statuses)
}
