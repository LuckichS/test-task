package repository

import (
	hezzl "Hezzl"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type ReprioritiizePostgres struct {
	db *sqlx.DB
	rd *redis.Client
}

func NewReprioritiizePostgres(db *sqlx.DB, rd *redis.Client) *ReprioritiizePostgres {
	return &ReprioritiizePostgres{db: db, rd: rd}
}

func (r *ReprioritiizePostgres) ReprioritiizeGood(good hezzl.Good) (int, error) {

	var oldPriority int
	queryOldPriority := fmt.Sprintf("SELECT priority FROM goods WHERE id=%d and project_id=%d", good.Id, good.Project_id)
	r.db.Get(&oldPriority, queryOldPriority)

	queryGetRows := fmt.Sprintf("SELECT id FROM goods WHERE priority>=%d;", oldPriority)

	var prioritiis []int
	err := r.db.Get(&prioritiis, queryGetRows)

	r.db.Exec(fmt.Sprintf(`UPDATE goods SET priority = %d WHERE id = %d;`, good.Priority, good.Id))
	//fmt.Println(fmt.Sprintf(`UPDATE goods SET priority = %d WHERE id = %d;`, oldPriority, good.Id))
	if err != nil {
		return 0, err
	}
	ctx := context.Background()
	r.rd.Del(ctx, fmt.Sprintf("%d", good.Id))
	return oldPriority, nil
}
