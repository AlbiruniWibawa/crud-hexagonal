package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// internal\adapter\repository\mongodb.go
type DB struct {
	*mongo.Database
}

func NewMongoDB(ctx context.Context, dsn, dbName string) (*DB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}
	return &DB{client.Database(dbName)}, nil
}
