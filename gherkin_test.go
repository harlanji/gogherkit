package gogherkit

import (
	"testing"
)

type LoginAPI struct {
  credentials map[string]string
}

func (loginAPI *LoginAPI) Init() {
  loginAPI.credentials = make(map[string]string, 100)
}

func (loginAPI *LoginAPI) Register(username string, password string) bool {
  if _, present := loginAPI.credentials[username]; present {
    return false
  }

  loginAPI.credentials[username] = password

  return true
}

func (loginAPI *LoginAPI) ChangePassword(username string, newPassword string) bool {
  if _, present := loginAPI.credentials[username]; !present {
    return false
  }

  loginAPI.credentials[username] = newPassword

  return true
}

func (loginAPI *LoginAPI) AttemptLogin(username string, password string) bool {
  if pw, present := loginAPI.credentials[username]; !present || pw != password {
  return false;
  }

  return true
}


func TestHandler(t *testing.T) {
	ggk := new(GoGherKit)

  var loginAPI *LoginAPI
  var result bool

  // BUG matchers with no tokens cause the system to fail
  ggk.AddMatcher("Given", "the login system with no $stuff", func(params StepFuncParam) {

    loginAPI = new(LoginAPI)
    loginAPI.Init()

  })

	ggk.AddMatcher("When", "I login with username $username and password $password", func(params StepFuncParam) {
		result = loginAPI.AttemptLogin(params["username"], params["password"])
	})
	ggk.AddMatcher("When", "I register with username $username and password $password", func(params StepFuncParam) {
		result = loginAPI.Register(params["username"], params["password"])
	})

	ggk.AddMatcher("Then", "the result should be $bool", func(params StepFuncParam) {
    if (result && params["bool"] != "true") || (!result && params["bool"] != "false") {
      t.Error("Expected the result to be ", params["bool"], "actual", result)
    }
  })


	ggk.RunFeatureFile("features/login.feature")
}
