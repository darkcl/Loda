package matcher

import "strings"

// MagnetMatcher will match a magnet link
type MagnetMatcher struct {
	Matcher
}

// Process will parse a string
func (u MagnetMatcher) Process(input string) bool {
	return strings.HasPrefix(input, "magnet:")
}

// Identifier describe matcher, used in getting downloader
func (u MagnetMatcher) Identifier() string {
	return "magnet"
}
