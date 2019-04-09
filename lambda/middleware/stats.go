package middleware

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/hamba/pkg/stats"
)

func WithRequestStats(h alexa.Handler, sable stats.Statable) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		if r.Type != alexa.TypeIntentRequest {
			h.Serve(b, r)
			return
		}

		tags := []interface{}{"intent", r.Intent.Name}

		stats.Inc(sable, "request.start", 1, 1.0, tags...)
		t := stats.Time(sable, "request.time", 1.0, tags...)

		h.Serve(b, r)

		t.Done()
		stats.Inc(sable, "request.complete", 1, 1.0, tags...)
	})
}
