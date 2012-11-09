package gogherkit

import (
  "testing"
  "fmt"
  "strings"
  "io/ioutil"
  "log"
)


type BeachCases struct {

  givenTodayIsHot StepFunc `blahb$la`

  whenTheTimeReaches StepFunc

  thenTheKidsAreAt StepFunc
}

func TestHandler(t *testing.T) {


	buffer, err := ioutil.ReadFile("features/beach.feature")
	if err != nil {
		log.Fatal(err)
	}

	gherkin := &Gherkin{Buffer: string(buffer)}
	gherkin.Init()

	if err := gherkin.Parse(); err != nil {
		log.Fatal(err)
	}



  AddMatcher("Given", "today it is $temp degrees $where", func(params StepFuncParam) {
    fmt.Printf("TODAY IT IS %s DEGREES %s!!!!\n", strings.ToUpper(params["temp"]), strings.ToUpper(params["where"]))
  })
  AddMatcher("When", "the time reaches $time", func(params StepFuncParam) {
    fmt.Printf("AND IT JUST TURNED %s!!!!\n", strings.ToUpper(params["time"]))
  })
  AddMatcher("Then", "the kids are at the $place", func(params StepFuncParam) {
    fmt.Printf("AND NOW THE KIDS ARE AT THE %s!!!!\n", strings.ToUpper(params["place"]))

  })


  gherkin.Execute()

}
