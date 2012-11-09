package gogherkit

import (
	//goLogger "github.com/googollee/go-logger"
	"io/ioutil"
)

type GoGherKit struct {
	StepManager
}

func (gogherkit *GoGherKit) RunFeatureFile(path string) {
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Err(err.Error())
	}

	gogherkit.RunFeatureText(string(buffer))
}

func (gogherkit *GoGherKit) RunFeatureText(content string) {
  gk := &Gherkin{Buffer: content}
	gk.StepManager = &gogherkit.StepManager
	gk.Init()

	if err := gk.Parse(); err != nil {
		logger.Err(err.Error())
	}

	gk.Execute()
}
