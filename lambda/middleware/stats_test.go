package middleware_test

import (
	"testing"
	"time"

	"github.com/drpsychick/alexa-go-cloudformation-demo/lambda/middleware"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/hamba/pkg/stats"
	"github.com/stretchr/testify/mock"
)

func TestWithRequestStats(t *testing.T) {
	tags := []interface{}{"locale", "en-US", "intent", "test-intent", "test-slot", "slot-value"}
	s := new(MockStats)
	s.On("Inc", "request.start", int64(1), float32(1.0), tags)
	s.On("Timing", "request.time", mock.Anything, float32(1.0), tags)
	s.On("Inc", "request.complete", int64(1), float32(1.0), tags)

	m := middleware.WithRequestStats(alexa.HandlerFunc(
		func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {

		}),
		stats.NewMockStatable(s),
	)

	bdr := &alexa.ResponseBuilder{}
	req := &alexa.RequestEnvelope{
		Request: &alexa.Request{
			Type: alexa.TypeIntentRequest,
			Intent: alexa.Intent{
				Name: "test-intent",
				Slots: map[string]*alexa.Slot{
					"test-slot": {
						Name:  "test-slot",
						Value: "slot-value",
					},
				},
			},
			Locale: "en-US",
		},
	}

	m.Serve(bdr, req)

	s.AssertExpectations(t)
}

func TestWithRequestStats_NonIntentRequests(t *testing.T) {
	tags := []interface{}{"locale", "en-US"}
	s := new(MockStats)
	s.On("Inc", "request.start", int64(1), float32(1.0), tags)
	s.On("Timing", "request.time", mock.Anything, float32(1.0), tags)
	s.On("Inc", "request.complete", int64(1), float32(1.0), tags)
	m := middleware.WithRequestStats(alexa.HandlerFunc(
		func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {

		}),
		stats.NewMockStatable(s),
	)

	bdr := &alexa.ResponseBuilder{}
	req := &alexa.RequestEnvelope{
		Request: &alexa.Request{
			Type:   alexa.TypeLaunchRequest,
			Locale: "en-US",
		},
	}

	m.Serve(bdr, req)

	s.AssertExpectations(t)
}

type MockStats struct {
	mock.Mock
}

func (m *MockStats) Inc(name string, value int64, rate float32, tags ...interface{}) {
	m.Called(name, value, rate, tags)
}

func (m *MockStats) Gauge(name string, value float64, rate float32, tags ...interface{}) {
	m.Called(name, value, rate, tags)
}

func (m *MockStats) Timing(name string, value time.Duration, rate float32, tags ...interface{}) {
	m.Called(name, value, rate, tags)
}

func (m *MockStats) Close() error {
	args := m.Called()
	return args.Error(0)
}
