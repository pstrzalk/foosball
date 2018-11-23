package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	result, err := Handler()
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	assert.IsType(t, nil, err)
}
