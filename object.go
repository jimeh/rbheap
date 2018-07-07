package main

import "encoding/json"

//go:generate easyjson -all object.go

func NewObject(inputJSON []byte) (*Object, error) {
	var obj Object
	err := json.Unmarshal(inputJSON, &obj)

	return &obj, err
}

type Object struct {
	Address string `json:"address"`
	Type    string `json:"type"`
}

func (s *Object) Index() string {
	return s.Address + ":" + s.Type
}
