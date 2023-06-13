package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	http.ListenAndServe(":8080", mux)
}
