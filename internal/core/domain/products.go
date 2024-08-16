package domain

import "time"

type Products struct {
	ID          uint64    `bson:"id"`
	ProductName string    `bson:"product_name"`
	Stock       uint64    `bson:"stock"`
	IsDeleted   bool      `bson:"is_deleted"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}
