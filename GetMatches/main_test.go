package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	request_string := "{\"per_page\":4,\"page\":1}"
	request := GetMatchesRequest{}

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
