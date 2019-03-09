package

import (
    "github.com/hamba/pkg/log"
    "github.com/aws/aws-lambda-go"
)

type Handler interface {
    Handle() (response, error)
}

type Application struct {
    logger  log.Logger
    Handler Handler
}

func NewApplication(l log.Logger) *Application {
    return &Application{
        logger: l,
    }
}

func (a *Application) Handle() (..., error) {
    response, err := a.Handler.Handle(r)
}

func (a *Application) Logger() log.Logger {
    return a.logger
}