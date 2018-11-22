package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	matches, err := Handler(Request{
		ID: 28,
	})
	assert.IsType(t, nil, err)
	assert.Equal(t, 2, len(matches))
}
