package service

import (
	test_app "github.com/AnnZh/test-app"
	"github.com/AnnZh/test-app/pkg/repository"
)

type QueriesService struct {
	repo repository.Queries
}

func NewQueriesService(repo repository.Queries) *QueriesService {
	return &QueriesService{repo: repo}
}

func (s *QueriesService) GetOverspeedCars(query test_app.OverSpeedQuery) ([]test_app.Message, error) {
	return s.repo.GetOverspeedCars(query)
}

func (s *QueriesService) GetMinMaxSpeedCars(query test_app.MinMaxQuery) ([]test_app.Message, error) {
	return s.repo.GetMinMaxSpeedCars(query)
}
