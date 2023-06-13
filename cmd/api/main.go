package main

import (
	"net/http"

	_ "github.com/frangar97/blogger/docs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title           Blogger API
// @version         1.0
func main() {
	mux := chi.NewRouter()

	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.Recoverer)
	mux.Use(handleCors)

	mux.Mount("/swagger", httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", mux)
}
