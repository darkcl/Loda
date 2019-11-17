package inspector_test

/*
This test file is to test on ci/cd environment
The aim of these tests is to ensure `youtube-dl` binary can be embbed within the application
*/

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/darkcl/loda/lib/inspector"
	"github.com/stretchr/testify/assert"
)

func binPath() string {
	dir, _ := os.Getwd()
	if runtime.GOOS == "windows" {
		return filepath.Join(dir, "../../bin/youtube-dl.exe")
	}

	return filepath.Join(dir, "../../bin/youtube-dl")
}

func TestInspectYoutubeURL(t *testing.T) {
	ytdl := inspector.NewYoutubeDLInspector(binPath())
	meta, err := ytdl.Process("https://www.youtube.com/watch?v=PC03Xgk__pg")
	assert.NotNil(t, meta)
	assert.Nil(t, err)

	data, _ := json.MarshalIndent(meta, "", "  ")
	t.Logf("Meta Data: %s\n", data)
}

func TestInspectNonYoutubeURL(t *testing.T) {
	ytdl := inspector.NewYoutubeDLInspector(binPath())
	meta, err := ytdl.Process("https://google.com")
	t.Logf("Error: %v\n", err)
	assert.Nil(t, meta)
	assert.NotNil(t, err)
}
