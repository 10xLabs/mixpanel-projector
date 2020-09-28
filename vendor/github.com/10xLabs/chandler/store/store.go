package store

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// Store ...
type Store interface {
	Connect()
	ReplaceOne(ctx context.Context, collectionName string, filter interface{}, p interface{}) (upsertedCount int64, err error)
	FetchOne(ctx context.Context, collectionName string, filter interface{}, p interface{}) error
	Fetch(ctx context.Context, collectionName string, filter interface{}, results interface{}) error
	FetchIDs(ctx context.Context, collectionName string, filter interface{}) (ids []string, err error)
	UpdateMany(ctx context.Context, collectionName string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (modifiedCount int64, err error)
	DeleteOne(ctx context.Context, collectionName string, filter interface{}) (deletedCount int64, err error)
	DeleteMany(ctx context.Context, collectionName string, filter interface{}) (deletedCount int64, err error)
	CountDocuments(ctx context.Context, collectionName string, filter interface{}) (int64, error)
}
