package repository

import (
	"context"

	"github.com/rinchsan/txdb-todo/pkg/entity"
)

//go:generate mockgen -source=./todo.go -destination=./mock/todo.go -package=mock -mock_names=Todo=Todo
type Todo interface {
	Add(ctx context.Context, todo *entity.Todo, assigneeUserIDs []uint64) error
	GetByID(ctx context.Context, id uint64) (*entity.Todo, error)
	GetAll(ctx context.Context) (entity.TodoSlice, error)
	Update(ctx context.Context, todo *entity.Todo, assigneeUserIDs []uint64) error
}
