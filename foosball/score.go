package foosball

import (
	"database/sql"
)

type Score struct {
	Id int `json:"id"`
	Players []Player `json:"players"`
	Score int `json:"score"`
}

func (s *Score) Save(db *sql.DB, m *Match) (int, error) {
	sql := `INSERT INTO scores (match_id, score) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(sql, m.Id, s.Score).Scan(&s.Id)
	if (err != nil) {
		return s.Id, err
	}

	for i := 0; i < 2; i++ {
		sql := `INSERT INTO player_scores (score_id, player_id) VALUES ($1, $2)`
		_, err := db.Exec(sql, s.Id, s.Players[i].Id)
		if (err != nil) {
			return 0, err
		}
	}

	return s.Id, err
}
