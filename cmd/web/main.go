package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/prranavv/reddit_clone/pkg/config"
	"github.com/prranavv/reddit_clone/pkg/driver"
	handler "github.com/prranavv/reddit_clone/pkg/handlers"
	"github.com/prranavv/reddit_clone/pkg/helpers"
	"github.com/prranavv/reddit_clone/pkg/render"
)

const port = "8080"

var session *scs.SessionManager
var app config.Appconfig
var logger *slog.Logger
var (
	Likedusers    = make(map[*string]bool)
	DisLikedusers = make(map[*string]bool)
)

func main() {
	//Create a logger
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	//Connecting the database
	db, err := driver.ConnectDB()
	if err != nil {
		logger.Error(
			err.Error(),
		)
		os.Exit(1)
	}
	logger.Info("Connected to Database")
	defer db.SQL.Close()
	//Loading the session and setting up constants
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	app.Session = session
	app.Logger = logger
	app.Likedusers = Likedusers
	app.DisLikedusers = DisLikedusers
	repo := handler.NewRepository(&app, db)
	handler.NewHandler(repo)
	helpers.NewHelpers(&app)
	render.NewRenderer(&app)

	srv := http.Server{
		Addr:    ":" + port,
		Handler: routes(),
	}
	logger.Info("Server is running on port 8080")
	err = srv.ListenAndServe()
	log.Fatal(err)
}
