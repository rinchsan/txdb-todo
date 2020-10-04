package user

import (
	"net/http"

	"github.com/rinchsan/txdb-todo/pkg/presenter"
)

func (h handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.user.GetAll(ctx)
	if err != nil {
		presenter.Error(w, err)
		return
	}

	presenter.Encode(w, users)
}
