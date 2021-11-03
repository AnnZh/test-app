package repository

import (
	"os"
	"strings"

	test_app "github.com/AnnZh/test-app"
)

type DataFile struct {
	f *os.File
}

func NewDataFile(f *os.File) *DataFile {
	return &DataFile{f: f}
}

func (r *DataFile) GetData(data test_app.Message) error {
	record := []string{data.Date, data.Number, data.Speed}
	mes := strings.Join(record, ";")

	n, err := r.f.Seek(0, os.SEEK_END)
	if err != nil {
		return err
	}

	if _, err := r.f.WriteAt([]byte(mes+"\n"), n); err != nil {
		return err
	}

	return nil
}
