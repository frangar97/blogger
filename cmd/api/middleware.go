package main

import "net/http"

func handleCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://*")

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT")
			w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Authorization")
		}

		next.ServeHTTP(w, r)
	})
}
