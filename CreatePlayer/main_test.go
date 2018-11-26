package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	request_string := "{\"name\":\"Sam Wise\"}"
	request := CreatePlayerRequest{}

	err := json.Unmarshal([]byte(request_string), &request)
	if err != nil {
		panic(err)
	}

	result, err := Handler(request)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	assert.IsType(t, nil, err)
}
