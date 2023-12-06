package main

import (
	"net/http"

	"github.com/prranavv/reddit_clone/pkg/helpers"
)

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsAuthenticated(r) {
			app.IsAuthenticated = 0
			next.ServeHTTP(w, r)
			return
		}
		app.IsAuthenticated = 1
		next.ServeHTTP(w, r)
	})
}

//The below code is what Tsawler had

// func Auth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if !helpers.IsAuthenticated(r) {
// 			session.Put(r.Context(), "error", "Please log in")
// 			http.Redirect(w, r, "/login", http.StatusSeeOther)
// 			return
// 		}
// 		app.IsAuthenticated = 1
// 		next.ServeHTTP(w, r)
// 	})
// }
