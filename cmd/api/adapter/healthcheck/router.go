package healthcheck

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rinchsan/txdb-todo/pkg/presenter"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	h := newHandler()

	r.Get("/", h.Get)

	return r
}

type handler struct {
}

func newHandler() handler {
	return handler{}
}

func (h handler) Get(w http.ResponseWriter, _ *http.Request) {
	presenter.Success(w)
}
