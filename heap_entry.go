package main

//go:generate easyjson -all heap_entry.go

// HeapEntry is a parsed heap item object
type HeapEntry struct {
	Address string `json:"address"`
	Type    string `json:"type"`
}
