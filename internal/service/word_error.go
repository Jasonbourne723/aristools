package service

import (
	"aristools/internal/dto"
	"encoding/json"
	"io"
	"math/rand"
	"os"
)

type errorWordService struct{}

var ErrorWordSrv = &errorWordService{}

func (s *errorWordService) Add(req []*dto.ErrorWordDto) error {
	dtos, err := s.read()
	if err != nil {
		return err
	}
	m := s.mapToDictionary(dtos)
	for _, item := range req {
		if word, exist := m[item.Id]; exist {
			word.Times++
		} else {
			dtos = append(dtos, item)
		}
	}
	return s.write(dtos)
}

func (s *errorWordService) Get(count int) ([]*dto.ErrorWordDto, error) {
	dtos, err := s.read()
	if err != nil {
		return nil, err
	}
	rand.Shuffle(len(dtos), func(i, j int) {
		dtos[i], dtos[j] = dtos[j], dtos[i]
	})
	if count > len(dtos) {
		count = len(dtos)
	}
	return dtos[:count], nil
}

func (s *errorWordService) Count() (int, error) {
	dtos, err := s.read()
	if err != nil {
		return 0, err
	}
	return len(dtos), err
}

func (s *errorWordService) mapToDictionary(words []*dto.ErrorWordDto) map[int64]*dto.ErrorWordDto {
	m := make(map[int64]*dto.ErrorWordDto, len(words))
	for _, item := range words {
		m[item.Id] = item
	}
	return m
}

func (s *errorWordService) read() ([]*dto.ErrorWordDto, error) {
	f, err := os.OpenFile(errorWordFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	words := []*dto.ErrorWordDto{}
	if err := json.NewDecoder(f).Decode(&words); err != nil {
		if err != io.EOF {
			return nil, err
		}
	}
	return words, nil
}

func (s *errorWordService) write(words []*dto.ErrorWordDto) error {
	f, err := os.OpenFile(errorWordFileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(&words)
}
