// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gogherkit

import (
	"io/ioutil"
	"log"
)

func main() {
	buffer, err := ioutil.ReadFile("doc/try.feature")
	if err != nil {
		log.Fatal(err)
	}

	gherkin := &Gherkin{Buffer: string(buffer)}
	gherkin.Init()

	if err := gherkin.Parse(); err != nil {
		log.Fatal(err)
	}
	//jBehave.Highlighter()

	gherkin.Execute()
}
