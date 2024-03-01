package repository

import (
	hezzl "Hezzl"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type RemovePostgres struct {
	db *sqlx.DB
	rd *redis.Client
}

func NewRemovePostgres(db *sqlx.DB, rd *redis.Client) *RemovePostgres {
	return &RemovePostgres{db: db, rd: rd}
}

func (r *RemovePostgres) RemoveGood(good hezzl.Good) error {
	queryCheck := fmt.Sprintf("SELECT id FROM goods WHERE id=%d AND project_id=%d", good.Id, good.Project_id)

	var isRow int
	err := r.db.Get(&isRow, queryCheck)

	if err != nil {
		return err
	}

	res, err := r.db.Exec(fmt.Sprintf(`DELETE FROM goods WHERE id = %d;`, good.Id))

	fmt.Println(res, err)
	ctx := context.Background()
	r.rd.Del(ctx, fmt.Sprintf("%d", good.Id))
	return nil
}
