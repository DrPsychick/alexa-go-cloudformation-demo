package alfalfa

import (
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"
)

type Application struct {
	logger  log.Logger
	statter stats.Statter
}

func NewApplication(l log.Logger, s stats.Statter) *Application {
	return &Application{
		logger:  l,
		statter: s,
	}
}

func (a *Application) Handle() {
	panic("implement me or panic hard")
}

func (a *Application) Help() (string, string) {
	return "", ""
}

func (a *Application) Logger() log.Logger {
	return a.logger
}

func (a *Application) Statter() stats.Statter {
	return a.statter
}
