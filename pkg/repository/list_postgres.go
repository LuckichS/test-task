package repository

import (
	hezzl "Hezzl"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type ListPostgres struct {
	db *sqlx.DB
	rd *redis.Client
}

func NewListPostgres(db *sqlx.DB, rd *redis.Client) *ListPostgres {
	return &ListPostgres{db: db, rd: rd}
}

func (r *ListPostgres) ListGood(good hezzl.Good, limit int, offset int) ([]hezzl.Good, int, error) {

	ctx := context.Background()
	key := fmt.Sprintf("%s-%s", limit, offset)
	//

	val, _ := r.rd.Get(ctx, key).Result()
	if len(val) > 0 {
		var res []hezzl.Good
		err := json.Unmarshal([]byte(val), &res)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			return res, 0, nil
		}
	}
	fmt.Println(val, "==")

	var goods []hezzl.Good
	queryToList := fmt.Sprintf("SELECT * FROM goods LIMIT %d OFFSET %d", limit, offset)
	err := r.db.Select(&goods, queryToList)
	if err != nil {
		return nil, 0, err
	}

	queryToRemoved := fmt.Sprintf("SELECT count(id) FROM goods WHERE removed=true LIMIT %d OFFSET %d ", limit, offset)
	var removed int

	err_ := r.db.Get(&removed, queryToRemoved)
	if err_ != nil {
		if err_.Error() != "sql: no rows in result set" {

			return nil, 0, err
		}
	}
	json_, _ := json.Marshal(goods)
	status, _ := r.rd.Set(ctx, key, json_, time.Duration(60*1000*1000*1000)).Result()
	fmt.Println(status, "--")
	return goods, removed, err

}
