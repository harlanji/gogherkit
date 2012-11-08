package gogherkit

import (
  "fmt"
)

type Story struct {
}

func (s *Story) BeginScenario(name string) {
  fmt.Printf("BEGIN SCENARIO: %s\n", name)
}

func (s *Story) EndScenario() {
  fmt.Println("END SCENARIO\n")
}
func (s *Story) BeginStory(name string) {
  fmt.Printf("BEGIN STORY: %s\n", name)
}
func (s *Story) EndStory() {
  fmt.Println("END STORY\n")
}
func (s *Story) StepType(buf string) {
  fmt.Printf("STEP TYPE: %s\n", buf)
}

func (s *Story) BeginStep(name string) {
  fmt.Printf("BEGIN STEP: %s\n", name)
}
func (s *Story) EndStep() {
  fmt.Println("END STEP\n")
}

type StepFunc func(map[string]string)




type stepMatcher struct {
  stepType string
  pattern string
  compiledPattern interface{}
  stepFunc StepFunc
}

var singleMatcher stepMatcher;

func FindStepMatcher(stepType string, sentence string) StepFunc {
  //sentence := "it is really $what"

  // match $_______ and fill an array with each as a key
  // turn ($_______) into a regex pattern, find matches
  // iterate matches, populating a second array that corresponds to the keys
  // pass the keys and values into the mapped function


  // we can probably pre-compile regex matchers for each step

  return singleMatcher.stepFunc

}


func AddMatcher(stepType string, pattern string, stepFunc StepFunc) {
  singleMatcher = stepMatcher{
    stepType: stepType,
    pattern: pattern,
    stepFunc: stepFunc,

    compiledPattern: nil, // reverse to a regex
  }
}
