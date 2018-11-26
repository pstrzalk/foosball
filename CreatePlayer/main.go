package main

import (
	"database/sql"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
	"github.com/pstrzalk/foosball/foosball"
)

var db *sql.DB

type CreatePlayerRequest struct {
	Name string `json:"name"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(r CreatePlayerRequest) (string, error) {
	db, err := foosball.InitDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	player := foosball.Player{Name: r.Name}

	_, err = player.Save(db)
	if err != nil {
		panic(err)
	}

	json_output, err := json.Marshal(player)
	if err != nil {
		panic(err)
	}

	return string(json_output), nil
}
