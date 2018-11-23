package main

import (
	"database/sql"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
	"github.com/pstrzalk/foosball/foosball"
)

var db *sql.DB

func main() {
	lambda.Start(Handler)
}

func Handler() (string, error) {
	players, err := getPlayers()
	if err != nil {
		return "", err
	}

	playersJson, err := json.Marshal(players)
	return string(playersJson), err
}

func getPlayers() ([]foosball.Player, error) {
  db, err := foosball.InitDb()
	if err != nil {
		return nil, err
	}
  defer db.Close()

	rows, err := db.Query("SELECT id, name FROM players")
	if err != nil {
		return nil, err
	}

	var players []foosball.Player
	var playerId int
	var playerName string

	for rows.Next() {
		err = rows.Scan(&playerId, &playerName)
		if err != nil {
			return nil, err
		}

		player := foosball.Player{
			Id:   playerId,
			Name: playerName,
		}
		players = append(players, player)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return players, nil
}
