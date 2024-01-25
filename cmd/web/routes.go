package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	handler "github.com/prranavv/reddit_clone/pkg/handlers"
)

func routes() *chi.Mux {
	//Initializing the router
	mux := chi.NewRouter()
	//middlewares
	mux.Use(SessionLoad)
	mux.Use(Auth)
	//Enables serving static files
	fileserver := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))
	//routes

	//login
	mux.Get("/login", handler.Repo.Login)
	mux.Post("/login", handler.Repo.PostLogin)
	mux.Get("/signup", handler.Repo.SignUp)
	mux.Post("/signup", handler.Repo.PostSignup)
	//logout
	mux.Get("/logout", handler.Repo.Logout)
	//home page
	mux.Get("/", handler.Repo.Home)
	//icon change
	mux.Post("/change-upicon/{post_id}/{logged_user}", handler.Repo.ChangeUpIcon)
	mux.Post("/change-downicon/{post_id}/{logged_user}", handler.Repo.ChangeDownIcon)
	//subreddit pages
	mux.Get("/r/askreddit", handler.Repo.Askreddit)
	mux.Get("/r/aww", handler.Repo.Aww)
	mux.Get("/r/funny", handler.Repo.Funny)
	mux.Get("/r/gaming", handler.Repo.Gaming)
	mux.Get("/r/todayilearned", handler.Repo.TodayIlearned)
	mux.Get("/r/worldnews", handler.Repo.Worldnews)
	mux.Get("/r/movies", handler.Repo.Movies)
	mux.Get("/r/food", handler.Repo.Food)
	//handling posting a post
	mux.Post("/create-post", handler.Repo.CreatePost)
	mux.Delete("/delete-post/{post_id}", handler.Repo.DeletePost)
	mux.Get("/getLikes/{post_id}", handler.Repo.GetLikesByPostID)
	return mux
}
