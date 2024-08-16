package http

import (
	"context"
	"strconv"

	"crud-hexagonal/internal/core/domain"
	"crud-hexagonal/internal/core/port"
	"crud-hexagonal/internal/core/utils"

	"github.com/gofiber/fiber/v2"
)

type ProductsHandler struct {
	service port.ProductsService
}

func NewProductsHandler(service port.ProductsService) *ProductsHandler {
	return &ProductsHandler{
		service,
	}
}

func (h *ProductsHandler) CreateProduct(c *fiber.Ctx) error {
	defer utils.TimeTrack("CreateProduct")()

	var req CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return handleError(c, 400, err)
	}

	product := domain.Products{
		ProductName: req.ProductName,
		Stock:       req.Stock,
	}

	newProduct, err := h.service.CreateProduct(context.Background(), &product)
	if err != nil {
		return handleError(c, 500, err)
	}

	handleResponse(c, 201, newProduct, "Product created successfully")

	return nil
}

func (h *ProductsHandler) GetProduct(c *fiber.Ctx) error {
	defer utils.TimeTrack("GetProduct")()

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return handleError(c, 400, err)
	}

	product, err := h.service.GetByID(context.Background(), id)
	if err != nil {
		return handleError(c, 500, err)
	}

	handleResponse(c, 200, product, "Product retrieved successfully")

	return nil
}

func (h *ProductsHandler) UpdateProduct(c *fiber.Ctx) error {
	defer utils.TimeTrack("UpdateProduct")()

	var req UpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return handleError(c, 400, err)
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return handleError(c, 400, err)
	}

	product := domain.Products{
		ID:          id,
		ProductName: req.ProductName,
		Stock:       req.Stock,
	}

	err = h.service.UpdateProduct(context.Background(), &product)
	if err != nil {
		return handleError(c, 500, err)
	}

	handleResponse(c, 200, product, "Product updated successfully")

	return nil
}

func (h *ProductsHandler) DeleteProduct(c *fiber.Ctx) error {
	defer utils.TimeTrack("DeleteProduct")()

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return handleError(c, 400, err)
	}

	err = h.service.DeleteProduct(context.Background(), id)
	if err != nil {
		return handleError(c, 500, err)
	}

	handleResponse(c, 200, nil, "Product deleted successfully")

	return nil
}

func (h *ProductsHandler) ListProducts(c *fiber.Ctx) error {
	defer utils.TimeTrack("ListProducts")()

	var productsList []domain.Products
	products, err := h.service.ListProducts(context.Background())
	if err != nil {
		return handleError(c, 500, err)
	}

	for _, product := range products {
		productsList = append(productsList, *product)
	}

	handleResponse(c, 200, productsList, "Products retrieved successfully")

	return nil
}
