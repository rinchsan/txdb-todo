package todo_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/rinchsan/txdb-todo/cmd/api/adapter/todo"
	"github.com/rinchsan/txdb-todo/pkg/entity"
	"github.com/rinchsan/txdb-todo/pkg/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetAll(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) todo.Handler
		code  int
	}{
		"repository.Todo.GetAll returns error": {
			setup: func(ctrl *gomock.Controller) todo.Handler {
				userRepo := mock.NewUser(ctrl)
				todoRepo := mock.NewTodo(ctrl)
				todoRepo.EXPECT().GetAll(gomock.Any()).Return(nil, errors.New("test error"))
				return todo.NewHandler(userRepo, todoRepo)
			},
			code: http.StatusInternalServerError,
		},
		"repository.Todo.GetAll succeeds": {
			setup: func(ctrl *gomock.Controller) todo.Handler {
				userRepo := mock.NewUser(ctrl)
				todoRepo := mock.NewTodo(ctrl)
				todoRepo.EXPECT().GetAll(gomock.Any()).Return(entity.TodoSlice{}, nil)
				return todo.NewHandler(userRepo, todoRepo)
			},
			code: http.StatusOK,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			h := c.setup(ctrl)

			r := httptest.NewRequest(http.MethodGet, "/todos", nil)
			w := httptest.NewRecorder()
			h.GetAll(w, r)
			assert.Equal(t, c.code, w.Code)
		})
	}
}
