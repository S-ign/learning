package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf generates cookie to prevent Cross-Site Scripting
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionsLoad loads and saves session cookies
func SessionsLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
