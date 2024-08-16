package service

import (
	"context"
	"time"

	"crud-hexagonal/internal/core/domain"
	"crud-hexagonal/internal/core/port"
)

type productsService struct {
	repo port.ProductsRepository
}

func NewProductsService(repo port.ProductsRepository) port.ProductsService {
	return &productsService{
		repo,
	}
}

func (s *productsService) CreateProduct(ctx context.Context, product *domain.Products) (*domain.Products, error) {
	product.IsDeleted = false
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	count, err := s.repo.CountProducts(ctx)
	if err != nil {
		return nil, err
	}
	product.ID = uint64(count + 1)

	return s.repo.CreateProducts(ctx, product)
}

func (s *productsService) GetByID(ctx context.Context, id uint64) (*domain.Products, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *productsService) UpdateProduct(ctx context.Context, product *domain.Products) error {
	product.UpdatedAt = time.Now()
	return s.repo.UpdateProducts(ctx, product)
}

func (s *productsService) DeleteProduct(ctx context.Context, id uint64) error {
	return s.repo.DeleteProducts(ctx, id)
}

func (s *productsService) ListProducts(ctx context.Context) ([]*domain.Products, error) {
	return s.repo.ListProducts(ctx)
}
