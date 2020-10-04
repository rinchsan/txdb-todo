package dao

import (
	"context"
	"database/sql"

	"github.com/rinchsan/txdb-todo/pkg/entity"
	"github.com/rinchsan/txdb-todo/pkg/repository"
	"github.com/rinchsan/txdb-todo/pkg/transaction"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func NewTodo(db *sql.DB) repository.Todo {
	return todoImpl{
		db: db,
	}
}

type todoImpl struct {
	db *sql.DB
}

func (impl todoImpl) Add(ctx context.Context, todo *entity.Todo, assigneeUserIDs []uint64) error {
	f := func(tx *sql.Tx) error {
		if err := todo.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}

		users, err := entity.Users(entity.UserWhere.ID.IN(assigneeUserIDs)).All(ctx, tx)
		if err != nil {
			return err
		}

		if err := todo.AddUsers(ctx, tx, false, users...); err != nil {
			return err
		}

		return nil
	}

	if err := transaction.Run(impl.db, f); err != nil {
		return err
	}

	return nil
}

func (impl todoImpl) GetByID(ctx context.Context, id uint64) (*entity.Todo, error) {
	todo, err := entity.FindTodo(ctx, impl.db, id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (impl todoImpl) GetAll(ctx context.Context) (entity.TodoSlice, error) {
	todos, err := entity.Todos(
		qm.Load(entity.TodoRels.AuthorUser),
		qm.Load(entity.TodoRels.Users),
	).All(ctx, impl.db)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (impl todoImpl) Update(ctx context.Context, todo *entity.Todo, assigneeUserIDs []uint64) error {
	f := func(tx *sql.Tx) error {
		if _, err := todo.Update(ctx, tx, boil.Infer()); err != nil {
			return err
		}

		users, err := entity.Users(entity.UserWhere.ID.IN(assigneeUserIDs)).All(ctx, tx)
		if err != nil {
			return err
		}

		if err := todo.SetUsers(ctx, tx, false, users...); err != nil {
			return err
		}

		return nil
	}

	if err := transaction.Run(impl.db, f); err != nil {
		return err
	}

	return nil
}
