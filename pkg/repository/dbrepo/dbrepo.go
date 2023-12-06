package dbrepo

import (
	"database/sql"

	"github.com/prranavv/reddit_clone/pkg/config"
	"github.com/prranavv/reddit_clone/pkg/repository"
)

type postgresRepo struct {
	App *config.Appconfig
	DB  *sql.DB
}

func NewPostgresRepo(app *config.Appconfig, db *sql.DB) repository.DBrepo {
	return &postgresRepo{
		App: app,
		DB:  db,
	}
}
