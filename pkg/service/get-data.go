package service

import (
	test_app "github.com/AnnZh/test-app"
	"github.com/AnnZh/test-app/pkg/repository"
)

type DataService struct {
	repo repository.Data
}

func NewDataService(repo repository.Data) *DataService {
	return &DataService{repo: repo}
}

func (s *DataService) GetData(data test_app.Message) error {
	return s.repo.GetData(data)
}
