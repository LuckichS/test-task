package service

import (
	hezzl "Hezzl"
	"Hezzl/pkg/repository"
	"time"
)

type CreateService struct {
	repo repository.Create
}

func NewCreateService(repo repository.Create) *CreateService {
	return &CreateService{repo: repo}
}

func (s CreateService) CreateGood(good hezzl.Good) (int, int, time.Time, error) {
	return s.repo.CreateGood(good)
}
