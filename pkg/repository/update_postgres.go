package repository

import (
	hezzl "Hezzl"
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type UpdatePostgres struct {
	db *sqlx.DB
	rd *redis.Client
}

func NewUpdatePostgres(db *sqlx.DB, rd *redis.Client) *UpdatePostgres {
	return &UpdatePostgres{db: db, rd: rd}
}

func (r *UpdatePostgres) UpdateGood(good hezzl.Good) (int, time.Time, error) {
	queryCheck := fmt.Sprintf("SELECT id FROM goods WHERE id=%d AND project_id=%d", good.Id, good.Project_id)
	now := time.Now()

	var isRow int
	err := r.db.Get(&isRow, queryCheck)

	if err != nil {
		return -1, now, err
	}

	r.db.Exec(fmt.Sprintf(`UPDATE goods SET name = '%s' WHERE id = %d;`, good.Name, good.Id))
	if len(good.Description) > 0 {
		r.db.Exec(fmt.Sprintf(`UPDATE goods SET description = '%s' WHERE id = %d;`, good.Description, good.Id))

	}
	ctx := context.Background()
	r.rd.Del(ctx, fmt.Sprintf("%d", good.Id))
	return 0, now, nil
}
