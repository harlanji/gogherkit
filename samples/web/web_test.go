package main

import (
	"fmt"
	"github.com/harlanji/gogherkit"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWeb(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler := CreateHttpHandler()

	ggk := new(gogherkit.GoGherKit)

	ggk.AddMatcher("Given", "the web server is $status", func(params gogherkit.StepFuncParam) {
		// todo: figure out how to start/stop it at will. for now, just assume we're good.
	})

	ggk.AddMatcher("When", "I request the main page with name '$name'", func(params gogherkit.StepFuncParam) {
		name := params["name"]

		req, err := http.NewRequest("GET", fmt.Sprintf("/%s", name), nil)

		if err != nil {
			t.Error("Could not create HTTP request")
		}

		handler.ServeHTTP(recorder, req)
	})

	ggk.AddMatcher("Then", "the text that comes back contains '$text'", func(params gogherkit.StepFuncParam) {
		body := params["text"]
		if !strings.Contains(string(recorder.Body.Bytes()), body) {
			t.Error("Body does not contain the expected text")
		}
	})

	ggk.RunFeatureFile("web.feature")

}
