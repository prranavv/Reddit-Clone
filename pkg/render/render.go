package render

import (
	"html/template"
	"net/http"

	"github.com/prranavv/reddit_clone/pkg/config"
	"github.com/prranavv/reddit_clone/pkg/models"
)

var app *config.Appconfig

func NewRenderer(a *config.Appconfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.IsAuthenticated = app.IsAuthenticated
	if td.IsAuthenticated == 1 {
		td.Username = app.Session.GetString(r.Context(), "username")
	}
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, templ string, templateData *models.TemplateData) {
	template, err := template.ParseFiles("./templates/"+templ, "./templates/login.layout.html", "./templates/base.layout.html", "./templates/subreddit.layout.html")
	if err != nil {
		app.Logger.Error(err.Error())
		return
	}
	td := AddDefaultData(templateData, r)
	err = template.Execute(w, td)
	if err != nil {
		app.Logger.Error(err.Error())
	}
}
