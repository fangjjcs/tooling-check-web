package main

import (
	"net/http"

	"github.com/fangjjcs/tooling/package/config"

	"github.com/go-chi/chi"

	"github.com/fangjjcs/tooling/package/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// mux.Use(middleware.Recoverer)
	// mux.Use(NoSurf)
	// mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/search", handlers.Repo.Search)
	mux.Post("/search", handlers.Repo.PostSearch) // It executes only when method is POST

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}


