package queryaws

type Application struct {
    Handler Handler
}

func New() *Application {
}

func (a *Application) Handle() (..., error) {
    response, err := a.Handler.Handle(r)

}


