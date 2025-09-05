package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"testing"
)

func TestDelete(t *testing.T) {
	storeOpts := StorageOptions{
		PathTransform: CASPathTransform,
	}

	store := NewStorage(storeOpts)

	key := "testdata"
	data := []byte("example data")

	err := store.WriteStream(key, bytes.NewReader(data))
	if err != nil {
		t.Error(err)
	}

	err = store.Delete(key)
	if err != nil {
		t.Error(err)
	}

	_, err = store.Read(key)
	if err == nil {
		t.Error("Expected error when reading deleted key, but got none")
	}
}

func TestTransformFunc(t *testing.T) {
	key := "my best pics"
	pathKey := CASPathTransform(key)
	log.Println("Transformed path:", pathKey)

	if pathKey.FileName != key {
		t.Errorf("Expected FileName %s but got %s", key, pathKey.FileName)
	}

	if pathKey.PathName == "" {
		t.Error("Expected non-empty PathName but got empty")
	}
	if len(pathKey.PathName) < 10 {
		t.Errorf("Expected PathName to be longer, got %s", pathKey.PathName)
	}
}

func TestStorage(t *testing.T) {
	store := newStorage()
	defer tearDown(store)

	for i := range 50 {
		key := fmt.Sprintf("testdata-%d", i)
		data := []byte("example data")

		err := store.WriteStream(key, bytes.NewReader(data))
		if err != nil {
			t.Error(err)
		}

		ok, err := store.Has(key)
		if err != nil {
			t.Errorf("Unexpected error in Has: %v", err)
		}
		if !ok {
			t.Error("Expected key to exist, but it does not")
		}

		readData, err := store.Read(key)
		if err != nil {
			t.Error(err)
		}

		b, _ := io.ReadAll(readData)
		if !bytes.Equal(b, data) {
			t.Errorf("Expected data %s but got %s", string(data), string(b))
		}

		if err := store.Delete(key); err != nil {
			t.Error(err)
		}

		ok, err = store.Has(key)
		if err != nil {
			t.Errorf("Unexpected error in Has after delete: %v", err)
		}
		if ok {
			t.Error("Expected key to be deleted, but it still exists")
		}
	}

}

func newStorage() *Storage {
	storeOpts := StorageOptions{
		PathTransform: CASPathTransform,
	}

	return NewStorage(storeOpts)
}

func tearDown(store *Storage) {
	err := store.clear()
	if err != nil {
		log.Println("Error during teardown:", err)
	}
}
