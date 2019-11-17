package downloader_test

import (
	"strconv"
	"testing"

	"github.com/darkcl/loda/lib/downloader"

	"github.com/stretchr/testify/assert"
)

func TestParsingDownloadProgress(t *testing.T) {
	sut := downloader.NewYoutubeDLDownloader(downloader.YoutubeDLDownloaderParams{}).(*downloader.YoutubeDLDownloader)
	progress := sut.ParseProgress(`[download]  67.6%% of 17.17MiB at  9.68MiB/s ETA 00:00`)
	f, _ := strconv.ParseFloat("67.6", 64)

	assert.Equal(t, f/100.0, progress.Progress)
}
