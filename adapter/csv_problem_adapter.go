package adapter

import (
	"encoding/csv"
	"errors"
	"io"
	"os"

	"github.com/with-shrey/go-quiz/domain"
)

var (
	ErrMalformedCsv = errors.New("csv row is malformed")
)

type CsvProblemAdapter struct {
	AddProblemService domain.AddProblemService
}

func (csvAdapter CsvProblemAdapter) ImportProblems(filePath string) error {
	reader, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer reader.Close()
	csvReader := csv.NewReader(reader)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if len(line) != 2 {
			return ErrMalformedCsv
		}
		csvAdapter.AddProblemService.AddProblem(line[0], line[1])
	}
	return nil
}
