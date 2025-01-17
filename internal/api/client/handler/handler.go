package handler

import (
	"context"
	"github.com/amzdll/shop/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, data model.ClientInfo) error
	GetByNameSurname(ctx context.Context, data model.ClientInfo) ([]model.Client, error)
	GetAll(ctx context.Context, pagination model.Pagination) ([]model.Client, error)
	UpdateClient(ctx context.Context, data model.Client) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type Handler struct {
	service   Service
	validator *validator.Validate
}

func New(s Service, v *validator.Validate) *Handler {
	return &Handler{
		service:   s,
		validator: v,
	}
}

func (h *Handler) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Post("/clients", h.Create)
	r.Get("/clients", h.GetAll)
	r.Get("/client", h.GetByNameSurname)
	r.Put("/clients", h.UpdateAddress)
	r.Delete("/clients/{id}", h.DeleteById)

	return r
}
