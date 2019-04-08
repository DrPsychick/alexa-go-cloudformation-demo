package alexa

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/json-iterator/go"
)

// Handler represents an alexa request handler.
type Handler interface {
	Serve(*Request) (*Response, error)
}

// HandlerFunc is an adapter allowing a function to be used as a handler.
type HandlerFunc func(*Request) (*Response, error)

// Serve serves the request.
func (fn HandlerFunc) Serve(r *Request) (*Response, error) {
	return fn(r)
}

// A Server defines parameters for running an Alexa server.
type Server struct {
	Handler Handler
}

// Invoke calls the handler, and serializes the response.
func (s *Server) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	req := &Request{}
	if err := jsoniter.Unmarshal(payload, req); err != nil {
		return nil, err
	}

	resp, err := s.Handler.Serve(req)
	if err != nil {
		return nil, err
	}

	return jsoniter.Marshal(resp)

}

// Serve serves the handler.
func (s *Server) Serve() error {
	// TODO: decide if we want a DefaultServeMux
	if s.Handler == nil {
		return errors.New("alexa: cannot serve empty handler")
	}

	lambda.StartHandler(s)
	return nil
}
