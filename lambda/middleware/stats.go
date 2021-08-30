package middleware

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/hamba/pkg/stats"
	"strings"
)

// WithRequestStats adds counter and timing stats to intent requests
func WithRequestStats(h alexa.Handler, sable stats.Statable) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		tags := []string{"locale", r.RequestLocale()}

		if r.Request.Type == alexa.TypeIntentRequest {
			tags = append(tags, "intent", r.IntentName())
		}
		for _, s := range r.Slots() {
			if s.Value != "" {
				tags = append(tags, strings.ToLower(s.Name), string(s.Value))
			}
		}

		stats.Inc(sable, "request.start", 1, 1.0, tags...)
		t := stats.Time(sable, "request.time", 1.0, tags...)

		h.Serve(b, r)

		t.Done()
		stats.Inc(sable, "request.complete", 1, 1.0, tags...)
	})
}
