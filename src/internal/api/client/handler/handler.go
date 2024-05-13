package handler

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	Create(ctx context.Context, clientInfo model.ClientInfo) error
	GetAll(ctx context.Context, clientListParams model.ClientListParams) ([]model.Client, error)
}

type Handler struct {
	service   Service
	validator *validator.Validate
}

func NewHandler(service Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

func (h *Handler) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Post("/clients", h.Create)
	r.Get("/clients/{name}_{surname}", h.GetByNameSurname)
	r.Get("/clients", h.GetAll)
	r.Put("/clients", h.UpdateAddress)
	r.Delete("/clients", h.DeleteById)

	return r
}