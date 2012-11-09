package gogherkit

import (
  "fmt"
  "regexp"
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

  // This is a $desc1 sentence with $desc2 tokens.
  //
  // find tokens using this:
  // (\$[^ ]+)

  // replace range with this:
  // This is a (.+) sentence with (.+) tokens.

  // then the matches are passed into params

  // match $_______ and fill an array with each as a key
  // turn ($_______) into a regex pattern, find matches
  // iterate matches, populating a second array that corresponds to the keys
  // pass the keys and values into the mapped function


  // we can probably pre-compile regex matchers for each step

  return singleMatcher.stepFunc

}


func AddMatcher(stepType string, pattern string, stepFunc StepFunc) {
  patternRgx := regexp.MustCompile("(\\$([0-9A-Za-z_]+))") // for some reason :word: doesn't work

  tokens := patternRgx.FindAllStringIndex(pattern, -1)

  fmt.Printf("Found %d tokens in pattern: %s\n", len(tokens), pattern)

  var newPattern = "";
  var startOffset = 0;
  for i, pair := range tokens {
    newPattern = fmt.Sprint(newPattern, string(pattern[startOffset:pair[0]]), "(.+)")
    startOffset = pair[1]
    fmt.Printf("Token(%d) from [%d:%d] = %s\n", i, pair[0], pair[1], pattern[pair[0]:pair[1]])
  }

  fmt.Printf("new pattern: %s\n", newPattern);


  singleMatcher = stepMatcher{
    stepType: stepType,
    pattern: pattern,
    stepFunc: stepFunc,

    compiledPattern: nil, // reverse to a regex
  }
}
