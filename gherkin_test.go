package gogherkit

import (
	"fmt"
	"strings"
	"testing"
)

type BeachCases struct {
	givenTodayIsHot StepFunc `blahb$la`

	whenTheTimeReaches StepFunc

	thenTheKidsAreAt StepFunc
}

func TestHandler(t *testing.T) {
  ggk := new(GoGherKit)

  ggk.LoadFeatureFile("features/beach.feature")

  /*
	ggk.AddSmartMatcher("Given today it is $temp degrees $where", func(temp String, where String) {
		fmt.Printf("TODAY IT IS %s DEGREES %s!!!!\n", strings.ToUpper(temp), strings.ToUpper(where))
	})
  */


	ggk.AddMatcher("Given", "today it is $temp degrees $where", func(params StepFuncParam) {
		fmt.Printf("TODAY IT IS %s DEGREES %s!!!!\n", strings.ToUpper(params["temp"]), strings.ToUpper(params["where"]))
	})

	ggk.AddMatcher("When", "the time reaches $time", func(params StepFuncParam) {
		fmt.Printf("AND IT JUST TURNED %s!!!!\n", strings.ToUpper(params["time"]))
	})

	ggk.AddMatcher("Then", "the kids are at the $place", func(params StepFuncParam) {
		fmt.Printf("AND NOW THE KIDS ARE AT THE %s!!!!\n", strings.ToUpper(params["place"]))
	})

	ggk.Run(t)
}
