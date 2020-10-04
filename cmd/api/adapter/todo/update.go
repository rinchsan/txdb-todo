package todo

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rinchsan/txdb-todo/pkg/presenter"
)

type UpdateInput struct {
	ID              uint64    `json:"id"`
	Title           string    `json:"title"`
	Detail          string    `json:"detail"`
	AuthorUserID    uint64    `json:"author_user_id"`
	DueDate         time.Time `json:"due_date"`
	AssigneeUserIDs []uint64  `json:"assignee_user_ids"`
}

func (h handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var in UpdateInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		presenter.BadRequest(w, err.Error())
		return
	}

	todo, err := h.todo.GetByID(ctx, in.ID)
	if err != nil {
		presenter.Error(w, err)
		return
	}

	todo.Title = in.Title
	todo.Detail = in.Detail
	todo.AuthorUserID = in.AuthorUserID
	todo.DueDate = in.DueDate
	if err := h.todo.Update(ctx, todo, in.AssigneeUserIDs); err != nil {
		presenter.Error(w, err)
		return
	}

	presenter.Success(w)
}
