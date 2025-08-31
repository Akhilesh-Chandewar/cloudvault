package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const defaultRootFolder = "root/network/storage"

func CASPathTransform(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashstr := hex.EncodeToString(hash[:])

	blockSize := 5
	sliceLength := len(hashstr) / blockSize
	dirs := make([]string, 0, blockSize)

	for i := range blockSize {
		start := i * sliceLength
		end := start + sliceLength
		if i == blockSize-1 {
			end = len(hashstr)
		}
		dirs = append(dirs, hashstr[start:end])
	}

	return PathKey{
		PathName: strings.Join(dirs, "/"),
		FileName: key,
	}
}

type PathTransform func(string) PathKey

type PathKey struct {
	PathName string
	FileName string
}

func (p PathKey) FullName() string {
	if p.PathName == "" {
		return p.FileName
	}
	return fmt.Sprintf("%s/%s", p.PathName, p.FileName)
}

type StorageOptions struct {
	Root          string
	PathTransform PathTransform
}

var DefaultPathTransform = func(key string) PathKey {
	return PathKey{PathName: "", FileName: key}
}

type Storage struct {
	options StorageOptions
}

func NewStorage(options StorageOptions) *Storage {
	if options.Root == "" {
		options.Root = defaultRootFolder
	}
	if options.PathTransform == nil {
		options.PathTransform = DefaultPathTransform
	}

	originalTransform := options.PathTransform
	options.PathTransform = func(key string) PathKey {
		pk := originalTransform(key)
		if options.Root != "" {
			if pk.PathName != "" {
				pk.PathName = fmt.Sprintf("%s/%s", options.Root, pk.PathName)
			} else {
				pk.PathName = options.Root
			}
		}
		return pk
	}

	return &Storage{
		options: options,
	}
}

func (s *Storage) Has(key string) (bool, error) {
	pathKey := s.options.PathTransform(key)
	_, err := os.Stat(pathKey.FullName())
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (s *Storage) clear() error {
	return os.RemoveAll(s.options.Root)
}

func (s *Storage) Delete(key string) error {
	pathKey := s.options.PathTransform(key)
	fullPath := pathKey.FullName()

	if err := os.RemoveAll(fullPath); err != nil {
		return err
	}

	dir := pathKey.PathName
	for dir != "" && dir != "." && dir != s.options.Root {
		err := os.Remove(dir)
		if err != nil {
			break
		}
		dir = filepath.Dir(dir)
	}

	log.Printf("Deleted %s and cleaned up empty parents", fullPath)
	return nil
}

func (s *Storage) Read(key string) (io.Reader, error) {
	f, err := s.ReadStream(key)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buff := new(bytes.Buffer)
	_, err = io.Copy(buff, f)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

func (s *Storage) ReadStream(key string) (io.ReadCloser, error) {
	pathKey := s.options.PathTransform(key)
	f, err := os.Open(pathKey.FullName())
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (s *Storage) Write(key string, r io.Reader) error {
	return s.WriteStream(key, r)
}

func (s *Storage) WriteStream(key string, r io.Reader) error {
	pathKey := s.options.PathTransform(key)
	if err := os.MkdirAll(pathKey.PathName, os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(pathKey.FullName())
	if err != nil {
		return err
	}
	defer f.Close()

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}

	fmt.Printf("wrote %d bytes to %s\n", n, pathKey.FullName())
	return nil
}
