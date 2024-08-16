package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"crud-hexagonal/internal/core/port"
)

func NewRouter(app *fiber.App, service port.ProductsService) {

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	RegisterRoutes(app, service)
}

func RegisterRoutes(app *fiber.App, service port.ProductsService) {
	handler := ProductsHandler{service}
	app.Post("/products", handler.CreateProduct)
	app.Get("/products/:id", handler.GetProduct)
	app.Put("/products/:id", handler.UpdateProduct)
	app.Delete("/products/:id", handler.DeleteProduct)
	app.Get("/products", handler.ListProducts)
}
