package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/rinchsan/txdb-todo/cmd/api/adapter/healthcheck"
	"github.com/rinchsan/txdb-todo/cmd/api/adapter/todo"
	"github.com/rinchsan/txdb-todo/cmd/api/adapter/user"
	"github.com/rinchsan/txdb-todo/pkg/logger"
	"github.com/rinchsan/txdb-todo/pkg/registry"
	"github.com/rinchsan/txdb-todo/pkg/server"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	flush, err := logger.Setup()
	if err != nil {
		panic(err)
	}
	defer flush()

	boil.DebugMode = true

	repo, cleanup := registry.NewRepository()
	defer cleanup()

	router := newRouter(repo)

	code := server.Run(router, 8080)
	os.Exit(code)
}

func newRouter(repo registry.Repository) http.Handler {
	r := chi.NewRouter()

	r.Mount("/", healthcheck.NewRouter())
	r.Mount("/users", user.NewRouter(repo))
	r.Mount("/todos", todo.NewRouter(repo))

	return r
}
