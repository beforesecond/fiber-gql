package api

import (
	"context"
	"log"
)

type TodoRepository interface {
	GetToDO(c context.Context, id string) string
}

type TodoRepo struct {
	TodoRepo TodoRepository
}

func (t *TodoRepo) GetToDO(c context.Context, id string) string {
	log.Print(id)
	return "555"
}
