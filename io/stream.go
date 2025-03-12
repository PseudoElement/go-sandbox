package iopack

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)

type PathTransformFunc func(string) PathKey

type PathKey struct {
	Pathname string
	Original string
}

var CASPathTransformFunc PathTransformFunc = func(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])

	blockSize := 5
	sliceLen := len(hashStr) / blockSize

	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blockSize, i*blockSize+blockSize
		paths[i] = hashStr[from:to]
	}

	return PathKey{
		Pathname: strings.Join(paths, "/"),
		Original: hashStr,
	}
}

var DefaultPathTransformFunc PathTransformFunc = func(key string) PathKey {
	return PathKey{key, key}
}

type StoreParams struct {
	PathTransformFunc PathTransformFunc
}

type Store struct {
	StoreParams
}

func NewStore(params StoreParams) *Store {
	return &Store{StoreParams: params}
}

func (s *Store) WriteStream(key string, r io.Reader) error {
	pathKey := s.PathTransformFunc(key)
	if err := os.MkdirAll(pathKey.Pathname, os.ModePerm); err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	io.Copy(buf, r)

	filenameBytes := md5.Sum(buf.Bytes())
	filename := hex.EncodeToString(filenameBytes[:])
	path := pathKey.Pathname + "/" + filename

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	n, err := io.Copy(f, buf)
	log.Printf("written %d bytes to disk: %s", n, path)

	return nil
}
