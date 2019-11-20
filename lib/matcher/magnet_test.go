package matcher_test

import (
	"testing"

	"github.com/darkcl/loda/lib/matcher"
	"github.com/stretchr/testify/assert"
)

func TestMatchMagnet(t *testing.T) {
	sut := &matcher.MagnetMatcher{}
	result := sut.Process(`magnet:?xt=urn:btih:546cf15f724d19c4319cc17b179d7e035f89c1f4&dn=ubuntu-14.04.2-desktop-amd64.iso&tr=http%%3A%%2F%%2Ftorrent.ubuntu.com%%3A6969%%2Fannounce&tr=http%%3A%%2F%%2Fipv6.torrent.ubuntu.com%%3A6969%%2Fannounce`)
	assert.True(t, result)
}

func TestMatchNonMagnet(t *testing.T) {
	sut := &matcher.MagnetMatcher{}
	result := sut.Process(`random string`)
	assert.False(t, result)
}
