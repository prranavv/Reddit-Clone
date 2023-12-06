package helpers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/prranavv/reddit_clone/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

var app *config.Appconfig

func NewHelpers(a *config.Appconfig) {
	app = a
}

func IsAuthenticated(r *http.Request) bool {
	exists := app.Session.Exists(r.Context(), "authenticated")
	return exists
}

func GenerateHashedPassword(password string) (string, error) {
	hashedpwd, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hashedpwd), err
}

func GettingSubRedditFromURL(r *http.Request) (string, error) {
	referer := r.Referer()
	parsedurl, err := url.Parse(referer)
	if err != nil {
		return "", err
	}
	fullpath := parsedurl.Path
	path := strings.Split(fullpath, "/")
	subreddit := path[2]
	return subreddit, nil
}
