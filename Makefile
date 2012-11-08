# Copyright 2010 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

gherkin: gherkin.peg.go gherkin.go main.go
	go build

gherkin.peg.go: gherkin.peg
	peg -switch -inline gherkin.peg

clean:
	rm -f gherkin gherkin.peg.go
