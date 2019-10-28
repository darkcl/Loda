package matcher

import "net/url"

// URLMatcher matches url (http, https only)
type URLMatcher struct {
	Matcher
}

// Process will parse a string an return next possible matcher
func (u URLMatcher) Process(input string) (bool, Matcher) {
	result, err := url.Parse(input)

	if err != nil {
		return false, nil
	}

	if result.Scheme != "http" && result.Scheme != "https" {
		return false, nil
	}

	return true, nil
}
