package main

import (
	"net/http"

	"github.com/S-ign/simple_webapp/pkg/config"
	handler "github.com/S-ign/simple_webapp/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionsLoad)

	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)

	return mux
}
