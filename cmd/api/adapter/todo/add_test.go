package todo_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/rinchsan/txdb-todo/cmd/api/adapter/todo"
	"github.com/rinchsan/txdb-todo/pkg/entity"
	"github.com/rinchsan/txdb-todo/pkg/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Add(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) todo.Handler
		body  string
		code  int
	}{
		"invalid json body": {
			setup: func(ctrl *gomock.Controller) todo.Handler {
				userRepo := mock.NewUser(ctrl)
				todoRepo := mock.NewTodo(ctrl)
				return todo.NewHandler(userRepo, todoRepo)
			},
			body: `{{}`,
			code: http.StatusBadRequest,
		},
		"repository.Todo.Add returns error": {
			setup: func(ctrl *gomock.Controller) todo.Handler {
				userRepo := mock.NewUser(ctrl)
				todoRepo := mock.NewTodo(ctrl)
				todoRepo.EXPECT().Add(gomock.Any(),
					&entity.Todo{
						Title:        "todo title",
						Detail:       "todo detail",
						AuthorUserID: 3,
						DueDate:      time.Date(2020, time.August, 20, 0, 0, 0, 0, time.UTC),
					},
					[]uint64{2, 3},
				).Return(errors.New("test error"))
				return todo.NewHandler(userRepo, todoRepo)
			},
			body: `{"title":"todo title", "detail":"todo detail", "due_date":"2020-08-20T00:00:00Z", "author_user_id":3, "assignee_user_ids":[2, 3]}`,
			code: http.StatusInternalServerError,
		},
		"repository.Todo.Add succeeds": {
			setup: func(ctrl *gomock.Controller) todo.Handler {
				userRepo := mock.NewUser(ctrl)
				todoRepo := mock.NewTodo(ctrl)
				todoRepo.EXPECT().Add(gomock.Any(),
					&entity.Todo{
						Title:        "todo title",
						Detail:       "todo detail",
						AuthorUserID: 3,
						DueDate:      time.Date(2020, time.August, 20, 0, 0, 0, 0, time.UTC),
					},
					[]uint64{2, 3},
				).Return(nil)
				return todo.NewHandler(userRepo, todoRepo)
			},
			body: `{"title":"todo title", "detail":"todo detail", "due_date":"2020-08-20T00:00:00Z", "author_user_id":3, "assignee_user_ids":[2, 3]}`,
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

			r := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(c.body))
			w := httptest.NewRecorder()
			h.Add(w, r)
			assert.Equal(t, c.code, w.Code)
		})
	}
}
