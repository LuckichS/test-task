package service

import (
	hezzl "Hezzl"
	"Hezzl/pkg/repository"
)

type ReprioritiizeService struct {
	repo repository.Reprioritiize
}

func NewReprioritiizeService(repo repository.Reprioritiize) *ReprioritiizeService {
	return &ReprioritiizeService{repo: repo}
}

func (s ReprioritiizeService) ReprioritiizeGood(good hezzl.Good) (int, error) {
	return s.repo.ReprioritiizeGood(good)
}
