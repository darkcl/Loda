package matcher

import "github.com/darkcl/loda/lib/inspector"

// YoutubeDLMatcher matches url with youtube-dl supported sites
type YoutubeDLMatcher struct {
	Matcher
	YtdlInspector inspector.Inspector
}

// NewYoutubeDLMatcher create a youtube-dl matcher
func NewYoutubeDLMatcher(binaryPath string) Matcher {
	return &YoutubeDLMatcher{
		YtdlInspector: inspector.NewYoutubeDLInspector(binaryPath),
	}
}

// Process will parse a string an return next possible matcher
func (u YoutubeDLMatcher) Process(input string) (bool, Matcher) {
	_, err := u.YtdlInspector.Process(input)
	if err != nil {
		return false, nil
	}

	return true, nil
}

// Identifier describe matcher, used in getting downloader
func (u YoutubeDLMatcher) Identifier() string {
	return "youtube-dl"
}
