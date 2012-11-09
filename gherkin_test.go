package gogherkit

import (
  "testing"
  "fmt"
  "strings"
  "io/ioutil"
  "log"
)

func TestHandler(t *testing.T) {


	buffer, err := ioutil.ReadFile("features/registration_and_login.feature")
	if err != nil {
		log.Fatal(err)
	}

	gherkin := &Gherkin{Buffer: string(buffer)}
	gherkin.Init()

	if err := gherkin.Parse(); err != nil {
		log.Fatal(err)
	}



  AddMatcher("Given", "today it is $temp degrees $where", func(params map[string]string) {
    fmt.Printf("TODAY IT IS %s DEGREES %s!!!!", strings.ToUpper(params["temp"]), strings.ToUpper(params["where"]))
  })
  AddMatcher("When", "the time reaches $time", func(params map[string]string) {
    fmt.Printf("AND IT JUST TURNED %s!!!!", strings.ToUpper(params["time"]))
  })
  AddMatcher("Then", "the kids are at the $place", func(params map[string]string) {
    fmt.Printf("AND NOW THE KIDS ARE AT THE %s!!!!", strings.ToUpper(params["place"]))
  })


  //gherkin.Execute()
}
