package gogherkit

import (
  "testing"
  "fmt"
  "strings"
  "io/ioutil"
  "log"
  "github.com/orfjackal/gospec/src/gospec"
)

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



  AddMatcher("Given", "today it is $temp degrees $where", func(params StepFuncParam, c *gospec.Context) {
    fmt.Printf("TODAY IT IS %s DEGREES %s!!!!\n", strings.ToUpper(params["temp"]), strings.ToUpper(params["where"]))
  })
  AddMatcher("When", "the time reaches $time", func(params StepFuncParam, c *gospec.Context) {
    fmt.Printf("AND IT JUST TURNED %s!!!!\n", strings.ToUpper(params["time"]))
  })
  AddMatcher("Then", "the kids are at the $place", func(params StepFuncParam, c *gospec.Context) {
    fmt.Printf("AND NOW THE KIDS ARE AT THE %s!!!!\n", strings.ToUpper(params["place"]))

    (*c).Expect(56, gospec.Equals, 56)
    (*c).Expect(56, gospec.Equals, 56)
    (*c).Expect(56, gospec.Equals, 56)
    (*c).Expect(56, gospec.Equals, 56)
  })


  gherkin.Execute()

  gospec.MainGoTest(gherkin.r, t)
}
