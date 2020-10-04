package user_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/rinchsan/txdb-todo/cmd/api/adapter/user"
	"github.com/rinchsan/txdb-todo/pkg/entity"
	"github.com/rinchsan/txdb-todo/pkg/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Update(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) user.Handler
		body  string
		code  int
	}{
		"invalid json body": {
			setup: func(ctrl *gomock.Controller) user.Handler {
				userRepo := mock.NewUser(ctrl)
				todoRepo := mock.NewTodo(ctrl)
				return user.NewHandler(userRepo, todoRepo)
			},
			body: `{{}`,
			code: http.StatusBadRequest,
		},
		"repository.User.GetByID returns error": {
			setup: func(ctrl *gomock.Controller) user.Handler {
				userRepo := mock.NewUser(ctrl)
				todoRepo := mock.NewTodo(ctrl)
				userRepo.EXPECT().GetByID(gomock.Any(), uint64(1)).Return(nil, errors.New("test error"))
				return user.NewHandler(userRepo, todoRepo)
			},
			body: `{"id":1, "username":"rinchsan"}`,
			code: http.StatusInternalServerError,
		},
		"repository.User.Update returns error": {
			setup: func(ctrl *gomock.Controller) user.Handler {
				userRepo := mock.NewUser(ctrl)
				todoRepo := mock.NewTodo(ctrl)
				userEntity := &entity.User{ID: 1}
				userRepo.EXPECT().GetByID(gomock.Any(), uint64(1)).Return(userEntity, nil)
				userEntity.Username = "rinchsan"
				userRepo.EXPECT().Update(gomock.Any(), userEntity).Return(errors.New("test error"))
				return user.NewHandler(userRepo, todoRepo)
			},
			body: `{"id":1, "username":"rinchsan"}`,
			code: http.StatusInternalServerError,
		},
		"repository.User.Update succeeds": {
			setup: func(ctrl *gomock.Controller) user.Handler {
				userRepo := mock.NewUser(ctrl)
				todoRepo := mock.NewTodo(ctrl)
				userEntity := &entity.User{ID: 1}
				userRepo.EXPECT().GetByID(gomock.Any(), uint64(1)).Return(userEntity, nil)
				userEntity.Username = "rinchsan"
				userRepo.EXPECT().Update(gomock.Any(), userEntity).Return(nil)
				return user.NewHandler(userRepo, todoRepo)
			},
			body: `{"id":1, "username":"rinchsan"}`,
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

			r := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(c.body))
			w := httptest.NewRecorder()
			h.Update(w, r)
			assert.Equal(t, c.code, w.Code)
		})
	}
}
