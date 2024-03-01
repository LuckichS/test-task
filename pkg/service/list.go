package service

import (
	hezzl "Hezzl"
	"Hezzl/pkg/repository"
)

type ListService struct {
	repo repository.List
}

func NewListService(repo repository.List) *ListService {
	return &ListService{repo: repo}
}

func (s ListService) ListGood(good hezzl.Good, limit int, offset int) ([]hezzl.Good, int, error) {

	return s.repo.ListGood(good, limit, offset)
}
