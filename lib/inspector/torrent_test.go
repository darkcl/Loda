package inspector_test

import (
	"testing"

	"github.com/darkcl/loda/lib/inspector"
	"github.com/stretchr/testify/assert"
)

func TestTorrentMetadata(t *testing.T) {
	sut := &inspector.TorrentInspector{}
	result, err := sut.Process("./testdata/test.torrent")
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestNonTorrentMetadata(t *testing.T) {
	sut := &inspector.TorrentInspector{}
	result, err := sut.Process("./testdata/test.txt")
	assert.Nil(t, result)
	assert.NotNil(t, err)
}
