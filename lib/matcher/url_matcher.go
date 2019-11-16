package matcher

import "net/url"

// URLMatcher matches url (http, https only)
type URLMatcher struct {
	Matcher
}

// Process will parse a string
func (u URLMatcher) Process(input string) bool {
	result, err := url.Parse(input)

	if err != nil {
		return false
	}

	if result.Scheme != "http" && result.Scheme != "https" {
		return false
	}

	return true
}

// Identifier describe matcher, used in getting downloader
func (u URLMatcher) Identifier() string {
	return "url"
}
