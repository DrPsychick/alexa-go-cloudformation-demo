package queryaws

import (
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/json-iterator/go"
)

type Application interface {
	QueryAWS(r interface{}) (interface{}, error)
}

// NewMux creates a new Mux instance.
func NewServer(app *Application) http.Handler {
	mux := bone.New()
	//mux.GetFunc("/test", requestHandlerTest(app))

	return mux
}

//func handleHelp(request alexa.Request) alexa.Response {
//	title := "Help"
//	SetLocale(request.Body.Locale)
//	r := alexa.NewSimpleTerminateResponse()
//	r.Body.OutputSpeech = &alexa.Payload{
//		Type:  OutputSpeechPlainText,
//		Text:  GetText("handle_help_response"),
//		Title: title,
//	}
//	r.Body.Reprompt = &alexa.Reprompt{
//		OutputSpeech: alexa.Payload{
//			Type: OutputSpeechPlainText,
//			Text: GetText("handle_help_reprompt"),
//		},
//	}
//	return r
//}

// WriteHtmlResponse encodes html content to the ResponseWriter.
func WriteJsonResponse(w http.ResponseWriter, code int, v interface{}) error {
	raw, err := jsoniter.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if _, err = w.Write(raw); err != nil {
		return err
	}

	return nil
}
