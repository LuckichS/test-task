package service

import (
	hezzl "Hezzl"
	"Hezzl/pkg/repository"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
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

type Service struct {
	Create
	Update
	Remove
	Reprioritiize
	List
	Redis *redis.Client
}

func NewService(repos *repository.Repository) *Service {
	redis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.db"),          // use default DB
	})

	return &Service{
		Create:        NewCreateService(repos.Create),
		Update:        NewUpdateService(repos.Update),
		Remove:        NewRemoveService(repos.Remove),
		Reprioritiize: NewReprioritiizeService(repos.Reprioritiize),
		List:          NewListService(repos.List),
		Redis:         redis,
	}
}
