package repository

import (
	"context"
	"time"

	"crud-hexagonal/internal/core/domain"
	"crud-hexagonal/internal/core/port"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoProductsRepository struct {
	db *mongo.Collection
}

func NewProductsRepository(db *mongo.Database) port.ProductsRepository {
	return &mongoProductsRepository{
		db: db.Collection("products"),
	}
}

func (r *mongoProductsRepository) CreateProducts(ctx context.Context, product *domain.Products) (*domain.Products, error) {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	_, err := r.db.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *mongoProductsRepository) GetByID(ctx context.Context, id uint64) (*domain.Products, error) {
	var product domain.Products
	filter := bson.M{"id": id}

	err := r.db.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *mongoProductsRepository) UpdateProducts(ctx context.Context, product *domain.Products) error {
	filter := bson.M{"id": product.ID}

	update := bson.M{
		"$set": product,
	}

	_, err := r.db.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoProductsRepository) DeleteProducts(ctx context.Context, id uint64) error {
	filter := bson.M{"id": id}

	update := bson.M{
		"$set": bson.M{
			"is_deleted": true,
			"updated_at": time.Now(),
		},
	}

	_, err := r.db.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoProductsRepository) ListProducts(ctx context.Context) ([]*domain.Products, error) {
	var products []*domain.Products
	cursor, err := r.db.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product domain.Products
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (r *mongoProductsRepository) CountProducts(ctx context.Context) (int64, error) {
	count, err := r.db.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}
	return count, nil
}
