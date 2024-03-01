package service

import (
	hezzl "Hezzl"
	"Hezzl/pkg/repository"
	"time"
)

type UpdateService struct {
	repo repository.Update
}

func NewUpdateService(repo repository.Update) *UpdateService {
	return &UpdateService{repo: repo}
}

func (s UpdateService) UpdateGood(good hezzl.Good) (int, time.Time, error) {
	return s.repo.UpdateGood(good)
}
