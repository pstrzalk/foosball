package main

import (
	// "encoding/json"
	// "errors"
	// "fmt"
	// "net/http"
	// "os"
	// "strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID int `json:"id"`
}

type Movie struct {
	Title       string `json:"title"`
	Description string `json:"overview"`
	Cover       string `json:"poster_path"`
	ReleaseDate string `json:"release_date"`
}

func Handler(request Request) ([]Movie, error) {
  movies := []Movie{
    Movie{
      Title: "Title1",
      Description: "Description1",
      Cover: "Cover1",
      ReleaseDate: "ReleaseDate1",
    },
    Movie{
      Title: "Title2",
      Description: "Description2",
      Cover: "Cover2",
      ReleaseDate: "ReleaseDate2",
    },
  }

	return movies, nil
}

func main() {
	lambda.Start(Handler)
}
