package main

import (
	"database/sql"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
	"github.com/pstrzalk/foosball/foosball"
)

var db *sql.DB

type CreateMatchScoreRequest struct {
	PlayerIds []int `json:"player_ids"`
	Score     int   `json:"score"`
}

type CreateMatchRequest struct {
	Scores []CreateMatchScoreRequest `json:"scores"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(r CreateMatchRequest) (string, error) {
	db, err := foosball.InitDb()
	if err != nil {
		return "", err
	}
	defer db.Close()

	players := []foosball.Player{
		foosball.Player{Id: r.Scores[0].PlayerIds[0]},
		foosball.Player{Id: r.Scores[0].PlayerIds[1]},
		foosball.Player{Id: r.Scores[1].PlayerIds[0]},
		foosball.Player{Id: r.Scores[1].PlayerIds[1]},
	}

	for i, player := range players {
		err = player.LoadById(db)
		if err != nil {
			return "", err
		}

		players[i] = player
	}

	match := foosball.Match{
		Scores: []foosball.Score{
			foosball.Score{
				Score:   r.Scores[0].Score,
				Players: players[:2],
			},
			foosball.Score{
				Score:   r.Scores[1].Score,
				Players: players[2:],
			},
		},
	}

	_, err = match.Save(db)
	if err != nil {
		return "", err
	}

	json_output, err := json.Marshal(match)
	if err != nil {
		return "", err
	}

	return string(json_output), nil
}
