package user

import (
	"encoding/json"
	"net/http"

	"github.com/rinchsan/txdb-todo/pkg/presenter"
)

type UpdateInput struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

func (h handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var in UpdateInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		presenter.BadRequest(w, err.Error())
		return
	}

	user, err := h.user.GetByID(ctx, in.ID)
	if err != nil {
		presenter.Error(w, err)
		return
	}

	user.Username = in.Username
	if err := h.user.Update(ctx, user); err != nil {
		presenter.Error(w, err)
		return
	}

	presenter.Success(w)
}
