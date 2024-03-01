package service

import (
	hezzl "Hezzl"
	"Hezzl/pkg/repository"
)

type RemoveService struct {
	repo repository.Remove
}

func NewRemoveService(repo repository.Remove) *RemoveService {
	return &RemoveService{repo: repo}
}

func (s RemoveService) RemoveGood(good hezzl.Good) error {
	return s.repo.RemoveGood(good)
}
