package matcher

import "path/filepath"

// TorrentMatcher will matcher a torrent file
type TorrentMatcher struct {
	Matcher
}

// Process will parse a string
func (u TorrentMatcher) Process(input string) bool {
	ext := filepath.Ext(input)
	return ext == ".torrent"
}

// Identifier describe matcher, used in getting downloader
func (u TorrentMatcher) Identifier() string {
	return "torrent"
}
