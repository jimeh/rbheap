package main

//go:generate easyjson -all heap_item.go

// HeapItem is a parsed heap item object
type HeapItem struct {
	Address string `json:"address"`
	Type    string `json:"type"`
}
