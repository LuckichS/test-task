package hezzl

import "time"

type Project struct {
	Id         int       `json:"-"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
}
