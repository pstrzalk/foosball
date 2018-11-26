package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	request_string := "{\"scores\":[{\"player_ids\":[1,2],\"score\":1},{\"player_ids\":[4,6],\"score\":3}]}"
	request := CreateMatchRequest{}

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
