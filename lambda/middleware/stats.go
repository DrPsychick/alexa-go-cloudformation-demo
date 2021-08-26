package middleware

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/lambda"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/hamba/pkg/stats"
	"strings"
)

// WithRequestStats adds counter and timing stats to intent requests
func WithRequestStats(h alexa.Handler, sable stats.Statable) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		tags := []interface{}{"locale", r.Locale}

		if r.Type == alexa.TypeIntentRequest {
			tags = append(tags, "intent", r.Intent.Name)
		}
		if len(r.Intent.Slots) > 0 {
			for _, s := range r.Intent.Slots {
				v := lambda.SlotValue(r, s.Name)
				if v != "" {
					tags = append(tags, strings.ToLower(s.Name), v)
				}
			}
		}

		stats.Inc(sable, "request.start", 1, 1.0, tags...)
		t := stats.Time(sable, "request.time", 1.0, tags...)

		h.Serve(b, r)

		t.Done()
		stats.Inc(sable, "request.complete", 1, 1.0, tags...)
	})
}
