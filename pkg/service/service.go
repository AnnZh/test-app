package service

import (
	test_app "github.com/AnnZh/test-app"
	"github.com/AnnZh/test-app/pkg/repository"
)

type Data interface {
	GetData(data test_app.Message) error
}

type Queries interface {
	GetOverspeedCars(query test_app.OverSpeedQuery) ([]test_app.Message, error)
	GetMinMaxSpeedCars(query test_app.MinMaxQuery) ([]test_app.Message, error)
}

type Service struct {
	Data
	Queries
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Data:    NewDataService(repos.Data),
		Queries: NewQueriesService(repos.Queries),
	}
}
