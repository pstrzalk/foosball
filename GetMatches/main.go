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

type GetMatchesRequest struct {
	PerPage int `json:"per_page"`
	Page    int `json:"page"`
}

func Handler(r GetMatchesRequest) (string, error) {
	matches, err := getMatches(r)
	if err != nil {
		return "", err
	}

	matchesJson, err := json.Marshal(matches)
	return string(matchesJson), err
}

func getMatches(r GetMatchesRequest) ([]foosball.Match, error) {
	db, err := foosball.InitDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if r.PerPage != 25 && r.PerPage != 50 {
		r.PerPage = 10
	}
	if r.Page < 1 {
		r.Page = 1
	}

	sqlStatement := `
		SELECT
		  m.id,
		  s1.id as score1_id,
		  s1.score as score1,
			s2.id as score2_id,
		  s2.score as score2,
		  p11.id as player11_id,
		  p12.id as player12_id,
		  p21.id as player21_id,
		  p22.id as player22_id,
		  p11.name as player11_name,
		  p12.name as player12_name,
		  p21.name as player21_name,
		  p22.name as player22_name
		FROM matches m
		LEFT JOIN scores s1 ON m.id = s1.match_id
		LEFT JOIN scores s2 ON m.id = s2.match_id AND s1.id <> s2.id

		LEFT JOIN player_scores ps11 ON s1.id = ps11.score_id
		LEFT JOIN player_scores ps12 ON s1.id = ps12.score_id
		LEFT JOIN player_scores ps21 ON s2.id = ps21.score_id
		LEFT JOIN player_scores ps22 ON s2.id = ps22.score_id

		LEFT JOIN players p11 ON p11.id = ps11.player_id
		LEFT JOIN players p12 ON p12.id = ps12.player_id
		LEFT JOIN players p21 ON p21.id = ps21.player_id
		LEFT JOIN players p22 ON p22.id = ps22.player_id

		WHERE
			p11.id <> p12.id AND
			p11.id <> p21.id AND
			p11.id <> p22.id AND
			p12.id <> p21.id AND
			p12.id <> p22.id AND
			p21.id <> p22.id AND
			s1.id < s2.id AND
			p11.id < p12.id AND
			p21.id < p22.id
		LIMIT $1 OFFSET $2
	`
	offset := r.PerPage * (r.Page - 1)
	rows, err := db.Query(sqlStatement, r.PerPage, offset)
	if err != nil {
		return nil, err
	}

	var matches []foosball.Match
	var matchId int
	var playerIds [4]int
	var playerNames [4]string
	var scoreIds [2]int
	var scores [2]int

	for rows.Next() {
		err = rows.Scan(
			&matchId,
			&scoreIds[0],
			&scores[0],
			&scoreIds[1],
			&scores[1],
			&playerIds[0],
			&playerIds[1],
			&playerIds[2],
			&playerIds[3],
			&playerNames[0],
			&playerNames[1],
			&playerNames[2],
			&playerNames[3],
		)
		if err != nil {
			return nil, err
		}

		match := foosball.Match{
			Id: matchId,
			Scores: []foosball.Score{
				foosball.Score{
					Id:    scoreIds[0],
					Score: scores[0],
					Players: []foosball.Player{
						foosball.Player{
							Id:   playerIds[0],
							Name: playerNames[0],
						},
						foosball.Player{
							Id:   playerIds[1],
							Name: playerNames[1],
						},
					},
				},
				foosball.Score{
					Id:    scoreIds[1],
					Score: scores[1],
					Players: []foosball.Player{
						foosball.Player{
							Id:   playerIds[2],
							Name: playerNames[2],
						},
						foosball.Player{
							Id:   playerIds[3],
							Name: playerNames[3],
						},
					},
				},
			},
		}
		matches = append(matches, match)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return matches, nil
}
