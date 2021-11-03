package repository

import (
	"bufio"
	"io"
	"math"
	"os"
	"strings"
	"sync"

	test_app "github.com/AnnZh/test-app"
)

func NewFile(name string) (*os.File, error) {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func GetMessagesByDate(f *os.File, date string) ([]test_app.Message, error) {
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	res, err := Process(f, date)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func Process(f *os.File, date string) ([]test_app.Message, error) {
	var result []test_app.Message

	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, 250*1024)
		return lines
	}}

	stringPool := sync.Pool{New: func() interface{} {
		lines := ""
		return lines
	}}

	r := bufio.NewReader(f)

	ch := make(chan []test_app.Message, 1)
	var wg sync.WaitGroup

	for {
		buf := linesPool.Get().([]byte)

		n, err := r.Read(buf)
		buf = buf[:n]

		if n == 0 {
			if err != nil {
				break
			}
			if err == io.EOF {
				break
			}
			return nil, err
		}

		nextUntillNewline, err := r.ReadBytes('\n')

		if err != io.EOF {
			buf = append(buf, nextUntillNewline...)
		}

		wg.Add(1)
		go func(ch chan<- []test_app.Message) {
			defer wg.Done()
			r := ProcessChunk(buf, &linesPool, &stringPool, date)
			ch <- r
		}(ch)
	}

	wg.Wait()
	close(ch)

	for res := range ch {
		result = append(result, res...)
	}
	return result, nil
}

func ProcessChunk(chunk []byte, linesPool *sync.Pool, stringPool *sync.Pool, date string) []test_app.Message {
	var result []test_app.Message

	var wg2 sync.WaitGroup
	ch := make(chan []test_app.Message, 1)

	mess := stringPool.Get().(string)
	mess = string(chunk)

	linesPool.Put(chunk)

	messSlice := strings.Split(mess, "\n")

	stringPool.Put(mess)

	chunkSize := 300
	n := len(messSlice)
	noOfThread := n / chunkSize

	if n%chunkSize != 0 {
		noOfThread++
	}

	for i := 0; i < (noOfThread); i++ {

		wg2.Add(1)
		go func(s int, e int, ch chan<- []test_app.Message, wg2 *sync.WaitGroup) {
			var res []test_app.Message

			defer wg2.Done()
			for i := s; i < e; i++ {
				text := messSlice[i]
				if len(text) == 0 {
					continue
				}
				mesSlice := strings.SplitN(text, ";", 2)
				mesCreationTimeString := mesSlice[0]

				if strings.Contains(mesCreationTimeString, date) {
					textSlice := strings.Split(text, ";")
					res = append(res, test_app.Message{Date: textSlice[0], Number: textSlice[1], Speed: textSlice[2]})
				}
			}
			ch <- res

		}(i*chunkSize, int(math.Min(float64((i+1)*chunkSize), float64(len(messSlice)))), ch, &wg2)
	}
	wg2.Wait()
	close(ch)

	for r := range ch {
		result = append(result, r...)
	}

	messSlice = nil

	return result
}
