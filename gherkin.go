package gogherkit

import (
	//goLogger "github.com/googollee/go-logger"
  "testing"
  "io/ioutil"
)


type GoGherKit struct {
  gk *Gherkin

  StepManager
}



func (gogherkit *GoGherKit) LoadFeatureFile(path string) {
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Err(err.Error())
	}

  gogherkit.LoadFeatureText(string(buffer))
}


func (gogherkit *GoGherKit) LoadFeatureText(content string) {
  gogherkit.gk = &Gherkin{Buffer: content}
  gogherkit.gk.StepManager = &gogherkit.StepManager
	gogherkit.gk.Init()

	if err := gogherkit.gk.Parse(); err != nil {
		logger.Err(err.Error())
	}
}



func (gogherkit *GoGherKit) Run(t *testing.T) {
  gogherkit.gk.Execute()
}

