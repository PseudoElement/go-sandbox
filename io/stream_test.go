package iopack

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathname := CASPathTransformFunc(key)
	fmt.Println("pathname", pathname)
}

func TestStore(t *testing.T) {
	params := StoreParams{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(params)

	data := bytes.NewReader([]byte("jpg file bytes"))
	if err := s.WriteStream("my-picture", data); err != nil {
		t.Error(err)
	}
}
