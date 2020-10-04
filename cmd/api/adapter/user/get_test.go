package user_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/rinchsan/txdb-todo/cmd/api/adapter/user"
	"github.com/rinchsan/txdb-todo/pkg/entity"
	"github.com/rinchsan/txdb-todo/pkg/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetAll(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) user.Handler
		code  int
	}{
		"repository.User.GetAll returns error": {
			setup: func(ctrl *gomock.Controller) user.Handler {
				userRepo := mock.NewUser(ctrl)
				todoRepo := mock.NewTodo(ctrl)
				userRepo.EXPECT().GetAll(gomock.Any()).Return(nil, errors.New("test error"))
				return user.NewHandler(userRepo, todoRepo)
			},
			code: http.StatusInternalServerError,
		},
		"repository.User.GetAll succeeds": {
			setup: func(ctrl *gomock.Controller) user.Handler {
				userRepo := mock.NewUser(ctrl)
				todoRepo := mock.NewTodo(ctrl)
				userRepo.EXPECT().GetAll(gomock.Any()).Return(entity.UserSlice{}, nil)
				return user.NewHandler(userRepo, todoRepo)
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

			r := httptest.NewRequest(http.MethodGet, "/users", nil)
			w := httptest.NewRecorder()
			h.GetAll(w, r)
			assert.Equal(t, c.code, w.Code)
		})
	}
}
