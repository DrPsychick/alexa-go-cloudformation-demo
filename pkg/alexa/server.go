package alexa

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/json-iterator/go"
	"sync"
)

// Handler represents an alexa request handler.
type Handler interface {
	Serve(*ResponseBuilder, *Request)
}

// HandlerFunc is an adapter allowing a function to be used as a handler.
type HandlerFunc func(*ResponseBuilder, *Request)

// Serve serves the request.
func (fn HandlerFunc) Serve(b *ResponseBuilder, r *Request) {
	fn(b, r)
}

// A Server defines parameters for running an Alexa server.
type Server struct {
	Handler Handler
}

// Invoke calls the handler, and serializes the response.
func (s *Server) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	req := &RequestEnvelope{}
	if err := jsoniter.Unmarshal(payload, req); err != nil {
		return nil, err
	}

	// TODO: panics when we get a wrong request
	req.Request.Context = req.Context
	req.Request.Session = req.Session

	builder := &ResponseBuilder{}
	s.Handler.Serve(builder, req.Request)

	return jsoniter.Marshal(builder.Build())

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

// Serve serves the given handler.
func Serve(h Handler) error {
	srv := &Server{Handler: h}

	return srv.Serve()
}

// ServeMux is an Alexa request multiplexer.
type ServeMux struct {
	mu           sync.RWMutex
	types        map[RequestType]Handler
	intents      map[string]Handler
	intent_slots map[string]string
}

// NewServeMux creates a new serve mux.
func NewServerMux() *ServeMux {
	return &ServeMux{
		types:        map[RequestType]Handler{},
		intents:      map[string]Handler{},
		intent_slots: map[string]string{},
	}
}

// Handler returns the matched handler for a request, or an error.
func (m *ServeMux) Handler(r *Request) (Handler, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if h, ok := m.types[r.Type]; ok {
		return h, nil
	}

	if r.Type != TypeIntentRequest {
		return nil, fmt.Errorf("server: unknown intent type %s", r.Type)
	}

	h, ok := m.intents[r.Intent.Name]
	if !ok {
		return nil, fmt.Errorf("server: unknown intent %s", r.Intent.Name)
	}

	return h, nil
}

// HandleRequestType registers the handler for the given request type.
//
// Any attempt to handle the IntentRequest type will be ignored, use Intent instead.
func (m *ServeMux) HandleRequestType(requestType RequestType, handler Handler) {
	if requestType == TypeIntentRequest {
		return
	}
	if handler == nil {
		panic("alexa: nil handler")
	}

	m.mu.Lock()

	m.types[requestType] = handler

	m.mu.Unlock()
}

// HandleRequestType registers the handler function for the given request type.
//
// Any attempt to handle the IntentRequest type will be ignored, use Intent instead.
func (m *ServeMux) HandleRequestTypeFunc(requestType RequestType, handler HandlerFunc) {
	m.HandleRequestType(requestType, handler)
}

// HandleIntent registers the handler for the given intent.
func (m *ServeMux) HandleIntent(intent string, handler Handler) {
	if handler == nil {
		panic("alexa: nil handler")
	}

	m.mu.Lock()

	m.intents[intent] = handler

	m.mu.Unlock()
}

// HandleIntent registers the handler function for the given intent.
func (m *ServeMux) HandleIntentFunc(intent string, handler HandlerFunc) {
	m.HandleIntent(intent, handler)
}

// Serve serves the matched handler.
func (m *ServeMux) Serve(b *ResponseBuilder, r *Request) {
	h, err := m.Handler(r)
	if err != nil {
		// TODO: Fallback handler
		return
	}

	h.Serve(b, r)
}

var DefaultServeMux = NewServerMux()

// HandleRequestType registers the handler for the given request type on the DefaultServeMux.
//
// Any attempt to handle the IntentRequest type will be ignored, use Intent instead.
func HandleRequestType(requestType RequestType, handler Handler) {
	DefaultServeMux.HandleRequestType(requestType, handler)
}

// HandleRequestType registers the handler function for the given request type on the DefaultServeMux.
//
// Any attempt to handle the IntentRequest type will be ignored, use Intent instead.
func HandleRequestTypeFunc(requestType RequestType, handler HandlerFunc) {
	DefaultServeMux.HandleRequestTypeFunc(requestType, handler)
}

// HandleIntent registers the handler for the given intent on the DefaultServeMux.
func HandleIntent(intent string, handler Handler) {
	DefaultServeMux.HandleIntent(intent, handler)
}

// HandleIntent registers the handler function for the given intent on the DefaultServeMux.
func HandleIntentFunc(intent string, handler HandlerFunc) {
	DefaultServeMux.HandleIntentFunc(intent, handler)
}
