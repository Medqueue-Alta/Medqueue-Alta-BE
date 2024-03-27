package services

import (
	"Medqueue-BE/features/todo"
	"Medqueue-BE/helper"
	"Medqueue-BE/middlewares"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	m todo.TodoModel
	v *validator.Validate
}

func NewTodoService(model todo.TodoModel) todo.TodoService {
	return &service{
		m: model,
		v: validator.New(),
	}
}

func (s *service) AddTodo(pemilik *jwt.Token, kegiatanBaru todo.Todo) (todo.Todo, error) {
	hp := middlewares.DecodeToken(pemilik)
	if hp == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return todo.Todo{}, errors.New("data tidak valid")
	}

	err := s.v.Struct(&kegiatanBaru)
	if err != nil {
		log.Println("error validasi", err.Error())
		return todo.Todo{}, err
	}

	result, err := s.m.InsertTodo(hp, kegiatanBaru)
	if err != nil {
		return todo.Todo{}, errors.New(helper.ServerGeneralError)
	}

	return result, nil
}
