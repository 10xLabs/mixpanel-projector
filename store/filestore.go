package store

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/10xLabs/chandler/store"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewFileStore ...
func NewFileStore() store.Store {
	return &FileStore{}
}

// FileStore ...
type FileStore struct{}

// ReplaceOne ...
func (f *FileStore) ReplaceOne(ctx context.Context, dir string, filter interface{}, p interface{}) (upsertedCount int64, err error) {
	data, err := json.Marshal(p)
	if err != nil {
		return 0, err
	}
	id := reflect.ValueOf(p).Elem().FieldByName("ID")

	err = ioutil.WriteFile(dir+"/"+id.String(), data, 0644)

	return 1, err
}

// FetchOne ...
func (f *FileStore) FetchOne(ctx context.Context, dir string, filter interface{}, p interface{}) error {
	id := reflect.ValueOf(p).Elem().FieldByName("ID")
	if !fileExists(dir + "/" + id.String()) {
		return nil
	}
	data, err := ioutil.ReadFile(dir + "/" + id.String())
	if err != nil {
		return err
	}

	return json.Unmarshal(data, p)
}

// Connect ...
func (f *FileStore) Connect() {
	panic("not implemented")
}

// Fetch ...
func (f *FileStore) Fetch(ctx context.Context, dir string, filter interface{}, results interface{}) error {
	panic("not implemented")
}

// FetchIDs ...
func (f *FileStore) FetchIDs(ctx context.Context, dir string, filter interface{}) (ids []string, err error) {
	panic("not implemented")
}

// UpdateMany ...
func (f *FileStore) UpdateMany(ctx context.Context, dir string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (modifiedCount int64, err error) {
	panic("not implemented")
}

// DeleteOne ...
func (f *FileStore) DeleteOne(ctx context.Context, dir string, filter interface{}) (deletedCount int64, err error) {
	panic("not implemented")
}

// DeleteMany ...
func (f *FileStore) DeleteMany(ctx context.Context, dir string, filter interface{}) (deletedCount int64, err error) {
	panic("not implemented")
}

// CountDocuments ...
func (f *FileStore) CountDocuments(ctx context.Context, dir string, filter interface{}) (int64, error) {
	panic("not implemented")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
