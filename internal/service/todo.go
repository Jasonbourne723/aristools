package service

import (
	"aristools/internal/dto"
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"
)

var TodoSrv = &TodoService{}

type TodoService struct{}

// 新增任务
func (s *TodoService) Add(req dto.AddTodoDto) error {

	todos, err := s.read()
	if err != nil {
		return err
	}

	for i := range todos {
		if todos[i].Name == req.Name {
			return errors.New("已存在同名任务")
		}
	}

	todos = append(todos, dto.TodoDto{
		Id:   s.getNewId(todos),
		Name: req.Name,
		DoAt: req.DoAt,
	})
	return s.write(todos)
}

func (s *TodoService) Del(ids []int64) error {
	if len(ids) == 0 {
		return nil
	}
	todos, err := s.read()
	if err != nil {
		return err
	}
	var newTodos []dto.TodoDto
	for _, item := range todos {
		isDeleted := false
		for _, id := range ids {
			if item.Id == id {
				isDeleted = true
			}
		}
		if !isDeleted {
			newTodos = append(newTodos, item)
		}
	}
	return s.write(newTodos)
}

func (s *TodoService) Today(ids []int64) error {

	if len(ids) == 0 {
		return nil
	}

	todos, err := s.read()
	if err != nil {
		return err
	}
	hasChanged := false
	for i, item := range todos {
		for _, id := range ids {
			if item.Id == id {
				todos[i].DoAt = time.Now().Format("2006-01-02")
				hasChanged = true
			}
		}
	}
	if hasChanged {
		s.write(todos)
	}
	return nil
}

func (s *TodoService) Done(ids []int64) error {
	todos, err := s.read()
	if err != nil {
		return err
	}
	for i, item := range todos {
		for _, id := range ids {
			if item.Id == id {
				todos[i].DoneAt = time.Now().Format("2006-01-02")
			}
		}
	}
	s.write(todos)
	return nil
}

// 列表
func (s *TodoService) List(today bool, all bool) ([]dto.TodoDto, error) {

	todos, err := s.read()
	if err != nil {
		return nil, err
	}
	if today {
		todays := make([]dto.TodoDto, 0)
		for _, item := range todos {
			if item.DoAt == time.Now().Format("2006-01-02") {
				todays = append(todays, item)
			}
		}
		return todays, nil
	}
	if all {
		return todos, nil
	}
	unDones := make([]dto.TodoDto, 0)
	for _, item := range todos {
		if len(item.DoneAt) == 0 {
			unDones = append(unDones, item)
		}
	}
	return unDones, nil
}

// 加载todos
func (s *TodoService) read() ([]dto.TodoDto, error) {

	todos := []dto.TodoDto{}
	filePath, err := getFilePath(todoFileName)
	if err != nil {
		return nil, err
	}
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	jsonDecoder := json.NewDecoder(f)
	if err := jsonDecoder.Decode(&todos); err != nil {
		if err != io.EOF {
			return nil, err
		}
	}

	return todos, nil
}

// 持久化todos
func (s *TodoService) write(todos []dto.TodoDto) error {

	filePath, err := getFilePath(todoFileName)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	jsonCoder := json.NewEncoder(f)
	if err := jsonCoder.Encode(&todos); err != nil {
		return err
	}
	return nil
}

// 获取新的Id
func (s *TodoService) getNewId(todos []dto.TodoDto) int64 {
	if len(todos) == 0 {
		return 1
	}
	return todos[len(todos)-1].Id + 1
}
