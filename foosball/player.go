package foosball

import (
	"database/sql"
)

type Player struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func (p *Player) Save(db *sql.DB) (int, error) {
  var err error

	if (p.Id > 0) {
		sqlStatement := `UPDATE players SET name = $2 WHERE id = $1`
		_, err = db.Exec(sqlStatement, p.Id, p.Name)
	} else {
		sqlStatement := `INSERT INTO players (name) VALUES ($1) RETURNING id`
		err = db.QueryRow(sqlStatement, p.Name).Scan(&p.Id)
	}

	return p.Id, err
}

func (p *Player) LoadById(db *sql.DB) error {
	sqlStatement := "SELECT name from players WHERE id = $1"
	row := db.QueryRow(sqlStatement, p.Id)

	return row.Scan(&p.Name)
}
