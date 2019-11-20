package matcher_test

import (
	"testing"

	"github.com/darkcl/loda/lib/matcher"
	"github.com/stretchr/testify/assert"
)

func TestMatchTorrent(t *testing.T) {
	sut := &matcher.TorrentMatcher{}
	result := sut.Process("/tmp/test.torrent")
	assert.True(t, result)
}

func TestMatchNonTorrent(t *testing.T) {
	sut := &matcher.TorrentMatcher{}
	result := sut.Process("/tmp/test.txt")
	assert.False(t, result)
}
