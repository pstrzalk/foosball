package foosball

import (
	"database/sql"
	"strconv"
)

type Match struct {
  Id int `json:"id"`
	Scores []Score `json:"scores"`
}

func (m *Match) Save(db *sql.DB) (int, error) {
	sql := `INSERT INTO matches (created_at) VALUES (NOW()) RETURNING id`
	err := db.QueryRow(sql).Scan(&m.Id)
	if (err != nil) {
		return m.Id, err
	}

	for i := 0; i < 2; i++ {
		_, err := m.Scores[i].Save(db, m)
		if (err != nil) {
			return 0, err
		}
	}

	return m.Id, err
}

func (m *Match) ToString() string {
	return strconv.Itoa(m.Scores[0].Players[0].Id) + "+" +
		strconv.Itoa(m.Scores[0].Players[1].Id) + "=" +
		strconv.Itoa(m.Scores[0].Score)
}
