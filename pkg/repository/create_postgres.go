package repository

import (
	hezzl "Hezzl"

	"time"

	"github.com/jmoiron/sqlx"
)

type CreatePostgres struct {
	db *sqlx.DB
}

func NewCreatePostgres(db *sqlx.DB) *CreatePostgres {
	return &CreatePostgres{db: db}
}

func (r *CreatePostgres) CreateGood(good hezzl.Good) (int, int, time.Time, error) {
	var now time.Time = time.Now().UTC()

	tx, err_ := r.db.Begin()
	if err_ != nil {
		return 0, 0, now, err_
	}

	var max_priority int
	queryToPririty := "SELECT MAX(priority) FROM goods"
	err := r.db.Get(&max_priority, queryToPririty)
	//fmt.Println(err.Error())
	if err != nil {
		max_priority = 0
	}

	var id int
	query := "INSERT INTO goods (project_id, name, description, priority, removed, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	row := tx.QueryRow(query, good.Project_id, good.Name, "", max_priority+1, false, now)
	if err := row.Scan(&id); err != nil {
		return 0, 0, now, err
	}

	return id, max_priority + 1, now, tx.Commit()
}
