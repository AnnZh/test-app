package repository

import (
	"os"
	"strconv"
	"strings"

	test_app "github.com/AnnZh/test-app"
)

type QueriesFile struct {
	f *os.File
}

func NewQueriesFile(f *os.File) *QueriesFile {
	return &QueriesFile{f: f}
}

func (r *QueriesFile) GetOverspeedCars(query test_app.OverSpeedQuery) ([]test_app.Message, error) {
	var result []test_app.Message

	res, err := GetMessagesByDate(r.f, query.Date)

	if err != nil {
		return nil, err
	}

	speed, err := strconv.ParseFloat(strings.Replace(query.Speed, ",", ".", 1), 64)

	if err != nil {
		return nil, err
	}

	for _, v := range res {
		curSpeed, err := strconv.ParseFloat(strings.Replace(v.Speed, ",", ".", 1), 64)

		if err != nil {
			return nil, err
		}

		if curSpeed > speed {
			result = append(result, v)
		}
	}

	return result, nil
}

type Mess struct {
	Date   string
	Number string
	Speed  float64
}

func (r *QueriesFile) GetMinMaxSpeedCars(query test_app.MinMaxQuery) ([]test_app.Message, error) {
	var result []test_app.Message
	var mess []Mess

	res, err := GetMessagesByDate(r.f, query.Date)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	for _, v := range res {
		curSpeed, err := strconv.ParseFloat(strings.Replace(v.Speed, ",", ".", 1), 64)

		if err != nil {
			return nil, err
		}

		mess = append(mess, Mess{Date: v.Date, Number: v.Number, Speed: curSpeed})
	}

	min := mess[0]
	max := mess[0]

	for _, v := range mess {
		if v.Speed > max.Speed {
			max = v
		}
		if v.Speed < min.Speed {
			min = v
		}
	}

	result = append(result, test_app.Message{
		Date:   max.Date,
		Number: max.Number,
		Speed:  strings.Replace(strconv.FormatFloat(max.Speed, 'f', 1, 64), ".", ",", 1),
	})
	result = append(result, test_app.Message{
		Date:   min.Date,
		Number: min.Number,
		Speed:  strings.Replace(strconv.FormatFloat(min.Speed, 'f', 1, 64), ".", ",", 1),
	})

	return result, nil
}
