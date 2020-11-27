package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/10xLabs/chandler/store"
	"github.com/10xLabs/chandler/tracer"
)

var tr = tracer.New()

// Repository ...
type Repository interface {
	Replace(ctx context.Context, id string, p interface{}) error
	Load(ctx context.Context, id string, p interface{}) error
	LoadByIDs(ctx context.Context, ids []string, r interface{}) error
	Delete(ctx context.Context, id string) error
}

// Base ...
type Base struct {
	Store          store.Store
	CollectionName *string
}

// New ...
func New(s store.Store, collectionName *string) Repository {
	return &Base{s, collectionName}
}

// Replace ...
func (b *Base) Replace(ctx context.Context, id string, p interface{}) (err error) {
	ts := tr.Trace(ctx, "repository.Replace")
	defer ts.Close(err)

	filter := bson.M{"_id": id}
	_, err = b.Store.ReplaceOne(ctx, *b.CollectionName, filter, p)

	return err
}

// Load ...
func (b *Base) Load(ctx context.Context, id string, p interface{}) (err error) {
	ts := tr.Trace(ctx, "repository.Load")
	defer ts.Close(err)

	filter := bson.M{"_id": id}
	err = b.Store.FetchOne(ctx, *b.CollectionName, filter, p)
	if err == store.ErrNoDocuments {
		return ErrNoDocuments
	}

	return err
}

// LoadByIDs ...
func (b *Base) LoadByIDs(ctx context.Context, ids []string, r interface{}) (err error) {
	ts := tr.Trace(ctx, "repository.LoadByIDs")
	defer ts.Close(err)

	filter := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}
	err = b.Store.Fetch(ctx, *b.CollectionName, filter, r)

	return err
}

// Delete ...
func (b *Base) Delete(ctx context.Context, id string) (err error) {
	ts := tr.Trace(ctx, "repository.Delete")
	defer ts.Close(err)

	filter := bson.M{"_id": id}
	_, err = b.Store.DeleteOne(ctx, *b.CollectionName, filter)

	return err
}
