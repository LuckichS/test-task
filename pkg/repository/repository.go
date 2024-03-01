package repository

import (
	hezzl "Hezzl"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Create interface {
	CreateGood(good hezzl.Good) (int, int, time.Time, error)
}

type Update interface {
	UpdateGood(good hezzl.Good) (int, time.Time, error)
}

type Remove interface {
	RemoveGood(good hezzl.Good) error
}

type Reprioritiize interface {
	ReprioritiizeGood(good hezzl.Good) (int, error)
}

type List interface {
	ListGood(good hezzl.Good, limit int, offset int) ([]hezzl.Good, int, error)
}

type Repository struct {
	Create
	Update
	Remove
	Reprioritiize
	List
}

func NewRepository(db *sqlx.DB, rd *redis.Client) *Repository {
	return &Repository{
		Create:        NewCreatePostgres(db),
		Update:        NewUpdatePostgres(db, rd),
		Remove:        NewRemovePostgres(db, rd),
		Reprioritiize: NewReprioritiizePostgres(db, rd),
		List:          NewListPostgres(db, rd),
	}
}
