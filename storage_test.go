package main

import (
	"bytes"
	"fmt"
	"io"
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
	fmt.Println("Transformed path:", pathKey)

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

	readData, err := store.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := io.ReadAll(readData)
	if !bytes.Equal(b, data) {
		t.Errorf("Expected data %s but got %s", string(data), string(b))
	}

	store.Delete(key)
}
