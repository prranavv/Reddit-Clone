package handler

import (
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/prranavv/reddit_clone/pkg/config"
	"github.com/prranavv/reddit_clone/pkg/driver"
	"github.com/prranavv/reddit_clone/pkg/helpers"
	"github.com/prranavv/reddit_clone/pkg/models"
	"github.com/prranavv/reddit_clone/pkg/render"
	"github.com/prranavv/reddit_clone/pkg/repository"
	"github.com/prranavv/reddit_clone/pkg/repository/dbrepo"
	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	App *config.Appconfig
	DB  repository.DBrepo
}

var Repo *Repository

func NewRepository(app *config.Appconfig, db *driver.DB) *Repository {
	return &Repository{
		App: app,
		DB:  dbrepo.NewPostgresRepo(app, db.SQL),
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

// Login handles the /login route
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	referer := r.Referer()
	parsedurl, err := url.Parse(referer)
	if err != nil {
		m.App.Logger.Error(
			err.Error(),
			slog.String("method", r.Method),
		)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	path := parsedurl.Path
	m.App.Session.Put(r.Context(), "path", path)
	render.RenderTemplate(w, r, "login.page.html", &models.TemplateData{})
}

// Logout handles the /login route
func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.Destroy(r.Context())
	_ = m.App.Session.RenewToken(r.Context())
	referer := r.Referer()
	parsedurl, err := url.Parse(referer)
	if err != nil {
		m.App.Logger.Error(
			err.Error(),
			slog.String("method", r.Method),
		)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	path := parsedurl.Path
	http.Redirect(w, r, path, http.StatusSeeOther)
}

func (m *Repository) ChangeUpIcon(w http.ResponseWriter, r *http.Request) {
	post_id := chi.URLParam(r, "post_id")
	int_post_id, err := strconv.Atoi(post_id)
	if err != nil {
		m.App.Logger.Error(err.Error())
		return
	}
	err = m.DB.AddingLikes(int_post_id)
	if err != nil {
		m.App.Logger.Error(err.Error(),
			slog.String("Type", "Database error"))
		return
	}
	no_of_likes, err := m.DB.GettingLikes(int_post_id)
	if err != nil {
		m.App.Logger.Error(err.Error(),
			slog.String("Type", "Database error"))
		return
	}
	render.RenderTemplate(w, r, "filled_up_icon.html", &models.TemplateData{No_Of_Likes: no_of_likes})

}

func (m *Repository) ChangeDownIcon(w http.ResponseWriter, r *http.Request) {
	post_id := chi.URLParam(r, "post_id")
	int_post_id, err := strconv.Atoi(post_id)
	if err != nil {
		m.App.Logger.Error(err.Error())
		return
	}
	err = m.DB.Disliking(int_post_id)
	if err != nil {
		m.App.Logger.Error(err.Error(),
			slog.String("Type", "Database error"))
		return
	}
	no_of_likes, err := m.DB.GettingLikes(int_post_id)
	if err != nil {
		m.App.Logger.Error(err.Error(),
			slog.String("Type", "Database error"))
		return
	}
	render.RenderTemplate(w, r, "filled_down_icon.html", &models.TemplateData{No_Of_Likes: no_of_likes})
}

// Signup handles the /signup route
func (m *Repository) SignUp(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "signup.page.html", &models.TemplateData{})
}

// Home renders the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

// PostLogin handles retreiving the data from the database and authenticating the user
func (m *Repository) PostLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	pwd_from_db, err := m.DB.GetPasswordFromUsername(username)
	if err != nil {
		m.App.Logger.Error(err.Error())
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(pwd_from_db), []byte(password))
	if err != nil {
		m.App.Logger.Error(err.Error())
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	path := m.App.Session.PopString(r.Context(), "path")
	m.App.Session.Put(r.Context(), "authenticated", "true")
	m.App.Session.Put(r.Context(), "username", username)
	http.Redirect(w, r, path, http.StatusSeeOther)
}

// PostSignup handles posting the signup form and writing the data to the database
func (m *Repository) PostSignup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	pwd, err := helpers.GenerateHashedPassword(password)
	if err != nil {
		m.App.Logger.Error(err.Error())
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}
	err = m.DB.InsertUserDetails(username, email, pwd)
	if err != nil {
		m.App.Logger.Error(err.Error())
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Askreddit renders the askreddit page
func (m *Repository) Askreddit(w http.ResponseWriter, r *http.Request) {
	fullpath := r.URL.Path
	path := strings.Split(fullpath, "/")
	subreddit := path[2]
	posts, err := m.DB.GetingPostsFromSubreddit(subreddit)
	if err != nil {
		m.App.Logger.Error(err.Error(),
			slog.String("method", r.Method),
			slog.String("Type", "Database Error"))
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = posts
	render.RenderTemplate(w, r, "askreddit.page.html", &models.TemplateData{Data: data})
}

// Aww renders the Aww page
func (m *Repository) Aww(w http.ResponseWriter, r *http.Request) {
	fullpath := r.URL.Path
	path := strings.Split(fullpath, "/")
	subreddit := path[2]
	posts, err := m.DB.GetingPostsFromSubreddit(subreddit)
	if err != nil {
		m.App.Logger.Error(err.Error(),
			slog.String("method", r.Method),
			slog.String("Type", "Database Error"))
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = posts
	render.RenderTemplate(w, r, "awwreddit.page.html", &models.TemplateData{Data: data})
}

// Funny renders the funny page
func (m *Repository) Funny(w http.ResponseWriter, r *http.Request) {
	fullpath := r.URL.Path
	path := strings.Split(fullpath, "/")
	subreddit := path[2]
	posts, err := m.DB.GetingPostsFromSubreddit(subreddit)
	if err != nil {
		m.App.Logger.Error(err.Error(),
			slog.String("method", r.Method),
			slog.String("Type", "Database Error"))
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = posts
	render.RenderTemplate(w, r, "funnyreddit.page.html", &models.TemplateData{Data: data})
}

// Gaming renders the gaming page
func (m *Repository) Gaming(w http.ResponseWriter, r *http.Request) {
	fullpath := r.URL.Path
	path := strings.Split(fullpath, "/")
	subreddit := path[2]
	posts, err := m.DB.GetingPostsFromSubreddit(subreddit)
	if err != nil {
		m.App.Logger.Error(err.Error(),
			slog.String("method", r.Method),
			slog.String("Type", "Database Error"))
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = posts
	render.RenderTemplate(w, r, "gamingreddit.page.html", &models.TemplateData{Data: data})
}

// Food renders the food page
func (m *Repository) Food(w http.ResponseWriter, r *http.Request) {
	fullpath := r.URL.Path
	path := strings.Split(fullpath, "/")
	subreddit := path[2]
	posts, err := m.DB.GetingPostsFromSubreddit(subreddit)
	if err != nil {
		m.App.Logger.Error(err.Error(),
			slog.String("method", r.Method),
			slog.String("Type", "Database Error"))
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = posts
	render.RenderTemplate(w, r, "foodreddit.page.html", &models.TemplateData{Data: data})
}

//Movies renders the Movies page

func (m *Repository) Movies(w http.ResponseWriter, r *http.Request) {
	fullpath := r.URL.Path
	path := strings.Split(fullpath, "/")
	subreddit := path[2]
	posts, err := m.DB.GetingPostsFromSubreddit(subreddit)
	if err != nil {
		m.App.Logger.Error(err.Error(),
			slog.String("method", r.Method),
			slog.String("Type", "Database Error"))
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = posts
	render.RenderTemplate(w, r, "moviesreddit.page.html", &models.TemplateData{Data: data})
}

// TodayIlearned renders the TodayIlearned page
func (m *Repository) TodayIlearned(w http.ResponseWriter, r *http.Request) {
	fullpath := r.URL.Path
	path := strings.Split(fullpath, "/")
	subreddit := path[2]
	posts, err := m.DB.GetingPostsFromSubreddit(subreddit)
	if err != nil {
		m.App.Logger.Error(err.Error(),
			slog.String("method", r.Method),
			slog.String("Type", "Database Error"))
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = posts
	render.RenderTemplate(w, r, "todayilearnedreddit.page.html", &models.TemplateData{Data: data})
}

// Worldnews renders the worldnews page
func (m *Repository) Worldnews(w http.ResponseWriter, r *http.Request) {
	fullpath := r.URL.Path
	path := strings.Split(fullpath, "/")
	subreddit := path[2]
	posts, err := m.DB.GetingPostsFromSubreddit(subreddit)
	if err != nil {
		m.App.Logger.Error(err.Error(),
			slog.String("method", r.Method),
			slog.String("Type", "Database Error"))
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = posts
	render.RenderTemplate(w, r, "worldnewsreddit.page.html", &models.TemplateData{Data: data})
}

type Post struct {
	Body  string
	Title string
}

// CreatePost creates a new post using htmx
func (m *Repository) CreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		m.App.Logger.Error(err.Error(), slog.String("method", r.Method))
		return
	}
	body := r.FormValue("body-text")
	title := r.FormValue("title-text")
	//adding the details to the stringmap
	stringmap := map[string]string{}
	stringmap["Body"] = body
	stringmap["Title"] = title
	if body == "" || title == "" {
		return
	}
	username := m.App.Session.Get(r.Context(), "username").(string)
	stringmap["Username"] = username
	//adding the details to the database
	subreddit, err := helpers.GettingSubRedditFromURL(r)
	if err != nil {
		m.App.Logger.Error(
			err.Error(),
			slog.String("method", r.Method),
		)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	post_ids, err := m.DB.CheckingDuplicatePost(body, title, subreddit, username)
	if err != nil {
		m.App.Logger.Error(err.Error(),
			slog.String("method", r.Method),
			slog.String("Type", "Database Error"))
		return
	}
	if len(post_ids) > 0 {
		m.App.Logger.Info("Duplicate post in database")
		return
	}
	err = m.DB.CreatePost(username, title, body, subreddit)
	if err != nil {
		m.App.Logger.Error(err.Error(),
			slog.String("method", r.Method),
			slog.String("Type", "Database Error"))
		return
	}
	post_id, err := m.DB.GettingPostIDFromDetails(username, title, body, subreddit)
	if err != nil {
		m.App.Logger.Error(
			err.Error(),
			slog.String("method", r.Method),
			slog.String("Type", "Database error"),
		)
		return
	}
	err = m.DB.InsertingDataIntoLikedTable(post_id)
	if err != nil {
		m.App.Logger.Error(
			err.Error(),
			slog.String("method", r.Method),
			slog.String("Type", "Database error"),
		)
		return
	}
	// posts, err := m.DB.GetingPostsFromSubreddit(subreddit)
	// if err != nil {
	// 	m.App.Logger.Error(err.Error(),
	// 		slog.String("method", r.Method),
	// 		slog.String("Type", "Database Error"))
	// 	return
	// }
	// log.Print(posts)
	// htmlstr := fmt.Sprintf("<br><div class='card'><br><div class='card-header'>u/PranavKumar</div><div class='card-body' ><blockquote class='blockquote mb-0'><p><strong>What is this post</strong></p><p>A well-known quote, contained in a blockquote element.</p><footer class='blockquote-footer'>Someone famous in <cite title='Source Title'>Source Title</cite></footer></blockquote></div></div>")
	// tmpl, _ := template.New("t").Parse(htmlstr)
	// tmpl.Execute(w, Post{Body: body, Title: title})
	render.RenderTemplate(w, r, "post.html", &models.TemplateData{StringMap: stringmap})
}
