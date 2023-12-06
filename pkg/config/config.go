package config

import (
	"log/slog"

	"github.com/alexedwards/scs/v2"
)

type Appconfig struct {
	Session         *scs.SessionManager
	IsAuthenticated int
	Logger          *slog.Logger
}
