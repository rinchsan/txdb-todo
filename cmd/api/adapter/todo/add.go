package todo

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rinchsan/txdb-todo/pkg/entity"
	"github.com/rinchsan/txdb-todo/pkg/presenter"
)

type AddInput struct {
	Title           string    `json:"title"`
	Detail          string    `json:"detail"`
	AuthorUserID    uint64    `json:"author_user_id"`
	DueDate         time.Time `json:"due_date"`
	AssigneeUserIDs []uint64  `json:"assignee_user_ids"`
}

func (in AddInput) Todo() entity.Todo {
	return entity.Todo{
		Title:        in.Title,
		Detail:       in.Detail,
		DueDate:      in.DueDate,
		AuthorUserID: in.AuthorUserID,
	}
}

func (h handler) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var in AddInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		presenter.BadRequest(w, err.Error())
		return
	}

	todo := in.Todo()
	if err := h.todo.Add(ctx, &todo, in.AssigneeUserIDs); err != nil {
		presenter.Error(w, err)
		return
	}

	presenter.Success(w)
}
