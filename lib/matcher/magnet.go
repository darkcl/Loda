package matcher

import "regexp"

// MagnetMatcher will match a magnet link
type MagnetMatcher struct {
	Matcher
}

// Process will parse a string
func (u MagnetMatcher) Process(input string) bool {
	magnetRegex := regexp.MustCompile(`^magnet:\?xt=urn:[a-z0-9]+:[a-z0-9]{32,40}&dn=.+&tr=.+$`)
	result := magnetRegex.MatchString(input)
	return result
}

// Identifier describe matcher, used in getting downloader
func (u MagnetMatcher) Identifier() string {
	return "magnet"
}
