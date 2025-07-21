package service

import (
	"aristools/internal/dto"
	"encoding/csv"
	"encoding/json"
	"io"
	"math/rand"
	"os"
	"strings"
)

type wordService struct{}

var WordSrv = &wordService{}

// 新增
func (s *wordService) Add(req dto.AddWordDto) error {
	words, err := s.read()
	if err != nil {
		return err
	}
	for _, item := range words {
		if strings.EqualFold(item.En, req.En) {
			return nil
		}
	}

	words = append(words, dto.WordDto{
		En: req.En,
		Cn: req.Cn,
		Id: s.getNewId(words),
	})
	return s.write(words)
}

// 统计数量
func (s *wordService) Count() (int64, error) {
	words, err := s.read()
	if err != nil {
		return 0, err
	}
	return int64(len(words)), nil
}

func (s *wordService) Rand(count int) ([]dto.WordDto, error) {
	words, err := s.read()
	if err != nil {
		return nil, err
	}
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	if count > len(words) {
		count = len(words)
	}
	return words[:count], nil
}

func (s *wordService) Import(filePath string) (int, error) {

	words, err := s.read()
	if err != nil {
		return 0, err
	}
	wordsMap := s.mapToDictionary(words)
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	csvReader.Comma = ','

	data, err := csvReader.ReadAll()
	if err != nil {
		return 0, err
	}
	var i int
	for _, item := range data {
		if _, exist := wordsMap[item[0]]; exist {
			continue
		}
		i++
		words = append(words, dto.WordDto{
			En: item[0],
			Cn: strings.Split(item[1], "、"),
			Id: s.getNewId(words),
		})
	}
	return i, s.write(words)
}

func (s wordService) read() ([]dto.WordDto, error) {
	f, err := os.OpenFile(wordFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	words := []dto.WordDto{}
	if err := json.NewDecoder(f).Decode(&words); err != nil {
		if err != io.EOF {
			return nil, err
		}
	}
	return words, nil
}

func (s *wordService) write(words []dto.WordDto) error {
	f, err := os.OpenFile(wordFileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(&words)
}

// 获取新的Id
func (s *wordService) getNewId(words []dto.WordDto) int64 {
	if len(words) == 0 {
		return 1
	}
	return words[len(words)-1].Id + 1
}

func (s *wordService) mapToDictionary(words []dto.WordDto) map[string]dto.WordDto {
	m := make(map[string]dto.WordDto, len(words))
	for _, item := range words {
		m[item.En] = item
	}
	return m
}
