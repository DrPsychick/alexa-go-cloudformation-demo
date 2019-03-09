package lambda

import (
	"context"

	"github.com/DrPsychick/alexa-go-cloudformation-demo/pkg/alexa"
)

type Application interface {
	QueryAWS(r interface{}) (interface{}, error)
}

func HandleRequest(ctx context.Context, r alexa.Request) (alexa.Response, error) {

}
