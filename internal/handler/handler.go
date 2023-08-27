package handler

import "github.com/gofiber/fiber/v2"

type Service interface{}

type Handler struct {
	S Service
}

func New(s Service) *Handler {
	return &Handler{
		S: s,
	}
}

func (h *Handler) GetApp() *fiber.App {
	app := fiber.New()

	return app
}
