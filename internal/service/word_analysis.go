package service

import (
	"aristools/internal/dto"
	"encoding/json"
	"io"
	"os"
	"time"
)

type wordAnalysisService struct{}

var WordAnalysisSrv = &wordAnalysisService{}

func (s *wordAnalysisService) Set(count int, errorCount int) error {

	dtos, err := s.read()
	if err != nil {
		return err
	}
	today := time.Now().Format("2006-01-02")

	if len(dtos) == 0 || dtos[len(dtos)-1].Date != today {
		dtos = append(dtos, &dto.WordAnalysisByDayDto{
			Date:     today,
			Count:    count,
			ErrCount: errorCount,
		})
	} else {
		dtos[len(dtos)-1].Count += count
		dtos[len(dtos)-1].ErrCount += errorCount
	}

	return s.write(dtos)
}

func (s *wordAnalysisService) GetToday() (*dto.WordAnalysisByDayDto, error) {
	dtos, err := s.read()
	if err != nil {
		return nil, err
	}
	today := time.Now().Format("2006-01-02")
	if len(dtos) > 0 && dtos[len(dtos)-1].Date == today {
		return dtos[len(dtos)-1], nil
	} else {
		return &dto.WordAnalysisByDayDto{
			Date:     today,
			Count:    0,
			ErrCount: 0,
		}, nil
	}

}

func (s *wordAnalysisService) GetAll() ([]*dto.WordAnalysisByDayDto, error) {
	dtos, err := s.read()
	if err != nil {
		return nil, err
	}
	for i, j := 0, len(dtos)-1; i < j; i, j = i+1, j-1 {
		dtos[i], dtos[j] = dtos[j], dtos[i]
	}
	return dtos, nil
}

func (s *wordAnalysisService) read() ([]*dto.WordAnalysisByDayDto, error) {
	f, err := os.OpenFile(wordAnalysisFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	words := []*dto.WordAnalysisByDayDto{}
	if err := json.NewDecoder(f).Decode(&words); err != nil {
		if err != io.EOF {
			return nil, err
		}
	}
	return words, nil
}

func (s *wordAnalysisService) write(words []*dto.WordAnalysisByDayDto) error {
	f, err := os.OpenFile(wordAnalysisFileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(&words)
}
