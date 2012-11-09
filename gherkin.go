package gogherkit

import (
	goLogger "github.com/googollee/go-logger"
)

type Story struct {
	stepType    string
	StepManager *StepManager
}

var logger, _ = goLogger.New(nil, "gogherkit.story")

func (s *Story) BeginScenario(name string) {
	logger.Debug("BEGIN SCENARIO: %s\n", name)
}

func (s *Story) EndScenario() {
	logger.Debug("END SCENARIO\n")
}
func (s *Story) BeginStory(name string) {
	logger.Debug("BEGIN STORY: %s\n", name)
}
func (s *Story) EndStory() {
	logger.Debug("END STORY\n")
}
func (s *Story) StepType(buf string) {
	logger.Debug("STEP TYPE: %s\n", buf)

	if buf != "And" {
		s.stepType = buf
	}
}

func (s *Story) BeginStep(name string) {
	logger.Debug("BEGIN STEP: %s\n", name)

	stepFunc, params := s.StepManager.FindStepMatcher(s.stepType, name)

	if stepFunc == nil {
		logger.Debug("Could not find step matcher for func %s\n", name)
		return
	}

	stepFunc(params)
}

func (s *Story) EndStep() {
	logger.Debug("END STEP\n")
}
