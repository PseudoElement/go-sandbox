package main

import "encoding/json"

type IBaseObject interface {
	MovementState() IMovementState
}

type IMovementState interface {
	CanMove() bool
	Move(dir string) error
	MoveImage() string
}

// =========================
type BaseMovementState struct {
}

func (ms *BaseMovementState) CanMove() bool {
	return false
}

func (ms *BaseMovementState) Move(dir string) error {
	return nil
}

func (ms *BaseMovementState) MoveImage() string {
	return ""
}

// =========================

type BaseObject struct {
	ms BaseMovementState
}

func (bo *BaseObject) MovementState() IMovementState {
	return &bo.ms
}

type NpcObject struct {
	BaseObject
}

func NewNpcObject() *NpcObject {
	return &NpcObject{}
}

// =========================

func checkNestedObjectAssignment() {
	s := []IBaseObject{NewNpcObject()}
	s = append(s, NewNpcObject())
}

func compareSlices[T any](s1, s2 []T) bool {
	buf1, err := json.Marshal(s1)
	if err != nil {
		panic(err)
	}
	buf2, err := json.Marshal(s2)
	if err != nil {
		panic(err)
	}
	if len(buf1) != len(buf2) {
		return false
	}
	for i := range len(buf1) {
		equal := buf1[i] == buf2[i]
		if !equal {
			return false
		}
	}
	return true
}
