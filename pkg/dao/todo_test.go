package dao_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rinchsan/txdb-todo/pkg/dao"
	"github.com/rinchsan/txdb-todo/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestTodoImpl_Add(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		todo            *entity.Todo
		assigneeUserIDs []uint64
		noErr           bool
	}{
		"new todo": {
			todo: &entity.Todo{
				Title:        "title",
				Detail:       "detail",
				DueDate:      time.Date(2020, time.December, 30, 0, 0, 0, 0, time.UTC),
				AuthorUserID: 1,
			},
			assigneeUserIDs: []uint64{1, 2},
			noErr:           true,
		},
		"empty assignee user ids": {
			todo: &entity.Todo{
				Title:        "title",
				Detail:       "detail",
				DueDate:      time.Date(2020, time.December, 30, 0, 0, 0, 0, time.UTC),
				AuthorUserID: 2,
			},
			assigneeUserIDs: []uint64{},
			noErr:           true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, err := sql.Open("txdb", uuid.New().String())
			assert.NoError(t, err)
			defer db.Close()
			impl := dao.NewTodo(db)

			err = impl.Add(context.Background(), c.todo, c.assigneeUserIDs)
			if c.noErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestTodoImpl_GetByID(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		id    uint64
		todo  *entity.Todo
		noErr bool
	}{
		"todo 1 exists": {
			id: 1,
			todo: &entity.Todo{
				ID:           1,
				Title:        "title 1",
				Detail:       "detail 1",
				AuthorUserID: 1,
				DueDate:      time.Date(2020, time.December, 30, 0, 0, 0, 0, time.UTC),
				CreatedAt:    time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
				UpdatedAt:    time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
			},
			noErr: true,
		},
		"todo 2 exists": {
			id: 2,
			todo: &entity.Todo{
				ID:           2,
				Title:        "title 2",
				Detail:       "detail 2",
				AuthorUserID: 2,
				DueDate:      time.Date(2020, time.December, 30, 0, 0, 0, 0, time.UTC),
				CreatedAt:    time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
				UpdatedAt:    time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
			},
			noErr: true,
		},
		"todo 100 does not exist": {
			id:    100,
			todo:  nil,
			noErr: false,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, err := sql.Open("txdb", uuid.New().String())
			assert.NoError(t, err)
			defer db.Close()
			impl := dao.NewTodo(db)

			todo, err := impl.GetByID(context.Background(), c.id)
			assert.Equal(t, c.todo, todo)
			if c.noErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestTodoImpl_GetAll(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		todos entity.TodoSlice
		noErr bool
	}{
		"all todos": {
			todos: entity.TodoSlice{
				&entity.Todo{
					ID:           1,
					Title:        "title 1",
					Detail:       "detail 1",
					AuthorUserID: 1,
					DueDate:      time.Date(2020, time.December, 30, 0, 0, 0, 0, time.UTC),
					CreatedAt:    time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
					UpdatedAt:    time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
				},
				&entity.Todo{
					ID:           2,
					Title:        "title 2",
					Detail:       "detail 2",
					AuthorUserID: 2,
					DueDate:      time.Date(2020, time.December, 30, 0, 0, 0, 0, time.UTC),
					CreatedAt:    time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
					UpdatedAt:    time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
				},
				&entity.Todo{
					ID:           3,
					Title:        "title 3",
					Detail:       "detail 3",
					AuthorUserID: 3,
					DueDate:      time.Date(2020, time.December, 30, 0, 0, 0, 0, time.UTC),
					CreatedAt:    time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
					UpdatedAt:    time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
				},
			},
			noErr: true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, err := sql.Open("txdb", uuid.New().String())
			assert.NoError(t, err)
			defer db.Close()
			impl := dao.NewTodo(db)

			todos, err := impl.GetAll(context.Background())
			for i := range todos {
				assert.Equal(t, c.todos[i].ID, todos[i].ID)
				assert.Equal(t, c.todos[i].Title, todos[i].Title)
				assert.Equal(t, c.todos[i].Detail, todos[i].Detail)
				assert.Equal(t, c.todos[i].AuthorUserID, todos[i].AuthorUserID)
				assert.Equal(t, c.todos[i].DueDate, todos[i].DueDate)
				assert.Equal(t, c.todos[i].CreatedAt, todos[i].CreatedAt)
				assert.Equal(t, c.todos[i].UpdatedAt, todos[i].UpdatedAt)
			}
			if c.noErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestTodoImpl_Update(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		todo            *entity.Todo
		assigneeUserIDs []uint64
		noErr           bool
	}{
		"todo 1": {
			todo: &entity.Todo{
				ID:           1,
				Title:        "new title",
				Detail:       "new detail",
				AuthorUserID: 1,
				DueDate:      time.Date(2020, time.December, 31, 0, 0, 0, 0, time.UTC),
			},
			assigneeUserIDs: []uint64{1, 2},
			noErr:           true,
		},
		"non-existing todo": {
			todo: &entity.Todo{
				ID:           100,
				Title:        "new title",
				Detail:       "new detail",
				AuthorUserID: 1,
				DueDate:      time.Date(2020, time.December, 31, 0, 0, 0, 0, time.UTC),
			},
			assigneeUserIDs: []uint64{1, 2},
			noErr:           false,
		},
		"empty title": {
			todo: &entity.Todo{
				ID:           3,
				Title:        "",
				Detail:       "new detail",
				AuthorUserID: 3,
				DueDate:      time.Date(2020, time.December, 31, 0, 0, 0, 0, time.UTC),
			},
			assigneeUserIDs: []uint64{1, 2},
			noErr:           true,
		},
		"empty assignee user ids": {
			todo: &entity.Todo{
				ID:           3,
				Title:        "new title",
				Detail:       "new detail",
				AuthorUserID: 3,
				DueDate:      time.Date(2020, time.December, 31, 0, 0, 0, 0, time.UTC),
			},
			assigneeUserIDs: []uint64{},
			noErr:           true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, err := sql.Open("txdb", uuid.New().String())
			assert.NoError(t, err)
			defer db.Close()
			impl := dao.NewTodo(db)

			err = impl.Update(context.Background(), c.todo, c.assigneeUserIDs)
			if c.noErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
