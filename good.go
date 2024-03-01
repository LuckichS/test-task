package hezzl

import "time"

type Good struct {
	Id          int       `json:"-"`
	Project_id  int       `json:"project_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Removed     bool      `json:"removed"`
	Created_at  time.Time `json:"created_at"` // binding:"required"`
}
