package main

import (
	"net/http"

	"github.com/S-ign/bookings/internal/config"
	handler "github.com/S-ign/bookings/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionsLoad)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	mux.Get("/contact", handler.Repo.Contact)
	mux.Get("/search-reservation", handler.Repo.SearchReservation)
	mux.Get("/make-reservation", handler.Repo.MakeReservation)
	mux.Get("/rooms/generals-quaters", handler.Repo.Generals)
	mux.Get("/rooms/majors-suite", handler.Repo.Majors)
	mux.Get("/reservation-summary", handler.Repo.ReservationSummary)

	mux.Post("/search-reservation", handler.Repo.PostSearchReservation)
	mux.Post("/search-availability-json", handler.Repo.SearchAvailabilityJSON)
	mux.Post("/make-reservation", handler.Repo.PostReservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
