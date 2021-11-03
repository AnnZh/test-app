package repository

import (
	"os"

	test_app "github.com/AnnZh/test-app"
)

type Data interface {
	GetData(data test_app.Message) error
}

type Queries interface {
	GetOverspeedCars(query test_app.OverSpeedQuery) ([]test_app.Message, error)
	GetMinMaxSpeedCars(query test_app.MinMaxQuery) ([]test_app.Message, error)
}

type Repository struct {
	Data
	Queries
}

func NewRepository(f *os.File) *Repository {
	return &Repository{
		Data:    NewDataFile(f),
		Queries: NewQueriesFile(f),
	}
}
