package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mayank550/bookings/pkg/config"
	"github.com/mayank550/bookings/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
    fileserver:=http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*",http.StripPrefix("/static",fileserver))
	return mux
}
