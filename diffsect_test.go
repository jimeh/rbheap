package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiffSect(t *testing.T) {
	tests := []struct {
		a      []string
		b      []string
		c      []string
		result []string
	}{
		{
			a:      []string{"1", "2", "3"},
			b:      []string{"1", "2", "3", "4", "5", "6"},
			c:      []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
			result: []string{"4", "5", "6"},
		},
		{
			a:      []string{"1", "2", "3"},
			b:      []string{"1", "2", "4", "5", "6"},
			c:      []string{"1", "5", "6", "7", "8", "9"},
			result: []string{"5", "6"},
		},
	}

	for _, pair := range tests {
		result := DiffSect(&pair.a, &pair.b, &pair.c)
		assert.Equal(t, *result, pair.result)
	}
}
