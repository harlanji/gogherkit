package main

import (
	"github.com/harlanji/gogherkit"
	"testing"
)

func TestHandler(t *testing.T) {
	ggk := new(gogherkit.GoGherKit)

	var loginAPI *LoginAPI
	var result bool

	ggk.AddMatcher("Given", "the login system with no data", func(params gogherkit.StepFuncParam) {

		loginAPI = new(LoginAPI)
		loginAPI.Init()

	})

	ggk.AddMatcher("When", "I login with username $username and password $password", func(params gogherkit.StepFuncParam) {
		result = loginAPI.AttemptLogin(params["username"], params["password"])
	})
	ggk.AddMatcher("When", "I register with username $username and password $password", func(params gogherkit.StepFuncParam) {
		result = loginAPI.Register(params["username"], params["password"])
	})

	ggk.AddMatcher("Then", "the result should be $bool", func(params gogherkit.StepFuncParam) {
		if (result && params["bool"] != "true") || (!result && params["bool"] != "false") {
			t.Error("Expected the result to be ", params["bool"], "actual", result)
		}
	})

	ggk.RunFeatureFile("login.feature")
}
