package store

import (
	"context"
	"time"

	"github.com/10xLabs/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoStore ...
type MongoStore struct {
	client       *mongo.Client
	database     *mongo.Database
	timeout      time.Duration
	hostURL      string
	databaseName string
}

// NewMongoStore ...
func NewMongoStore(mongoDBURL, mongoDBDatabase string) Store {
	return &MongoStore{
		hostURL:      mongoDBURL,
		databaseName: mongoDBDatabase,
		timeout:      10 * time.Second,
	}
}

func logOperation(operation string, filter interface{}) {
	log.WithFields(log.Fields{
		"operation": operation,
		"filter":    filter,
		"count":     1,
	}).Info("store operation")
}

// ReplaceOne ...
func (ms *MongoStore) ReplaceOne(ctx context.Context, collectionName string, filter interface{}, p interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, ms.timeout)
	defer cancel()

	collection := ms.database.Collection(collectionName)
	defer logOperation("ReplaceOne", filter)
	res, err := collection.ReplaceOne(ctx, filter, p, options.Replace().SetUpsert(true))
	if res == nil {
		return 0, err
	}

	return res.UpsertedCount, err
}

// FetchOne ...
func (ms *MongoStore) FetchOne(ctx context.Context, collectionName string, filter interface{}, p interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, ms.timeout)
	defer cancel()

	collection := ms.database.Collection(collectionName)
	defer logOperation("FindOne", filter)
	if err := collection.FindOne(ctx, filter).Decode(p); err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNoDocuments
		}
		return err
	}

	return nil
}

// FetchIDs ...
func (ms *MongoStore) FetchIDs(ctx context.Context, collectionName string, filter interface{}) (ids []string, err error) {
	ctx, cancel := context.WithTimeout(ctx, ms.timeout)
	defer cancel()

	collection := ms.database.Collection(collectionName)
	defer logOperation("Find", filter)
	cur, err := collection.Find(ctx, filter, options.Find().SetProjection(bson.D{}))
	if err != nil {
		return nil, err
	}

	ctx1, cancel1 := context.WithTimeout(ctx, ms.timeout)
	defer cancel1()

	for cur.Next(ctx1) {
		f := struct {
			ID string `bson:"_id"`
		}{}

		if err := cur.Decode(&f); err != nil {
			return nil, err
		}
		ids = append(ids, f.ID)
	}

	return ids, nil
}

// Fetch ...
func (ms *MongoStore) Fetch(ctx context.Context, collectionName string, filter interface{}, results interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, ms.timeout)
	defer cancel()

	collection := ms.database.Collection(collectionName)
	defer logOperation("Find", filter)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return err
	}

	ctx1, cancel1 := context.WithTimeout(ctx, ms.timeout)
	defer cancel1()

	return cur.All(ctx1, results)
}

// UpdateMany ...
func (ms *MongoStore) UpdateMany(ctx context.Context, collectionName string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, ms.timeout)
	defer cancel()

	collection := ms.database.Collection(collectionName)
	defer logOperation("UpdateMany", filter)
	res, err := collection.UpdateMany(ctx, filter, update, opts...)
	if res == nil {
		return 0, err
	}

	return res.ModifiedCount, err
}

// DeleteOne ...
func (ms *MongoStore) DeleteOne(ctx context.Context, collectionName string, filter interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, ms.timeout)
	defer cancel()

	collection := ms.database.Collection(collectionName)
	defer logOperation("DeleteOne", filter)
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}

// DeleteMany ...
func (ms *MongoStore) DeleteMany(ctx context.Context, collectionName string, filter interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, ms.timeout)
	defer cancel()

	collection := ms.database.Collection(collectionName)
	defer logOperation("DeleteMany", filter)
	res, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}

// CountDocuments ...
func (ms *MongoStore) CountDocuments(ctx context.Context, collectionName string, filter interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, ms.timeout)
	defer cancel()

	collection := ms.database.Collection(collectionName)
	defer logOperation("CountDocuments", filter)

	return collection.CountDocuments(ctx, filter)
}

// Connect ...
func (ms *MongoStore) Connect() {
	var err error
	ms.client, err = mongo.NewClient(options.Client().ApplyURI(ms.hostURL))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), ms.timeout)
	defer cancel()

	err = ms.client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	ctx, cancel = context.WithTimeout(ctx, ms.timeout)
	defer cancel()

	err = ms.client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	ms.database = ms.client.Database(ms.databaseName)
}
