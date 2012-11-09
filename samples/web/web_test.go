package main

import (
	"github.com/harlanji/gogherkit"
	"testing"
)

func TestWeb(t *testing.T) {
	RunWebServer()

	ggk := new(gogherkit.GoGherKit)

	ggk.AddMatcher("Given", "the web server is $status", func(params gogherkit.StepFuncParam) {
		// todo: figure out how to start/stop it at will. for now, just assume we're good.
	})

	ggk.AddMatcher("When", "I request the main page with name $name", func(params gogherkit.StepFuncParam) {
		//name := params["name"]

		// make web request to http://localhost:8080/ + name
		// store response where the Then can get it
	})

	ggk.AddMatcher("Then", "The text that comes back is '$text'", func(params gogherkit.StepFuncParam) {
		// use retrieved response from When, and verify contents.

	})

	ggk.RunFeatureFile("web.feature")

	StopWebServer()
}
