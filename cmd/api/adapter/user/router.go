package user

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rinchsan/txdb-todo/pkg/registry"
	"github.com/rinchsan/txdb-todo/pkg/repository"
)

func NewRouter(repo registry.Repository) http.Handler {
	r := chi.NewRouter()
	h := newHandler(repo)

	r.Get("/", h.GetAll)
	r.Post("/", h.Add)
	r.Put("/", h.Update)

	return r
}

type handler struct {
	user repository.User
	todo repository.Todo
}

func newHandler(repo registry.Repository) handler {
	return handler{
		user: repo.User(),
		todo: repo.Todo(),
	}
}
