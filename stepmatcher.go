package gogherkit

import (
	"fmt"
	"regexp"
)

type StepFuncParam map[string]string
type StepFunc func(StepFuncParam)

type stepMatcher struct {
	stepType        string
	pattern         string
	paramNames      []string
	compiledPattern *regexp.Regexp
	stepFunc        StepFunc
}

type StepManager struct {
	matchers []stepMatcher
}

func (sm StepManager) FindStepMatcher(stepType string, sentence string) (StepFunc, StepFuncParam) {
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

	for _, m := range sm.matchers {
		if m.stepType == stepType && m.compiledPattern.MatchString(sentence) {
			params := make(StepFuncParam, len(m.paramNames))

			mrng := m.compiledPattern.FindAllStringSubmatchIndex(sentence, -1)[0]

			for i, name := range m.paramNames {
				value := sentence[mrng[2*(i+1)]:mrng[2*(i+1)+1]]

				logger.Debug("Param(%d) %s = %s\n", i, name, value)

				params[name] = value
			}

			logger.Debug("applying. param %s = %s\n", m.paramNames[0], sentence[mrng[2]:mrng[3]])

			return m.stepFunc, params
		}
	}

	return nil, nil

}

func (sm *StepManager) AddMatcher(stepType string, pattern string, stepFunc StepFunc) {
	patternRgx := regexp.MustCompile("(\\$([0-9A-Za-z_]+))") // for some reason :word: doesn't work

	tokens := patternRgx.FindAllStringSubmatchIndex(pattern, -1)

	logger.Debug("Found %d tokens in pattern: %s\n", len(tokens), pattern)

	var newPattern = ""
	var startOffset = 0
	var params = make([]string, len(tokens))

	for i, pair := range tokens {
		params[i] = pattern[pair[4]:pair[5]]

		newPattern = fmt.Sprint(newPattern, string(pattern[startOffset:pair[0]]), "(.+)")
		startOffset = pair[1]
		logger.Debug("Param(%d) from [%d:%d], with identifier from [%d:%d] = %s\n", i, pair[0], pair[1], pair[4], pair[5], params[i])

	}

	logger.Debug("new pattern: %s\n", newPattern)

	matcher := stepMatcher{
		stepType: stepType,
		pattern:  pattern,
		stepFunc: stepFunc,

		paramNames:      params,
		compiledPattern: regexp.MustCompile(newPattern), // reverse to a regex
	}

	sm.matchers = append(sm.matchers, matcher)
}
