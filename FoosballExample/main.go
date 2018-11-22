package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID int `json:"id"`
}

type Match struct {
	Title string `json:"title"`
}

func Handler(request Request) ([]Match, error) {
  matches := []Match{
    Match{
      Title: "Title1",
    },
    Match{
      Title: "Title2",
    },
  }

	return matches, nil
}

func main() {
	lambda.Start(Handler)
}
