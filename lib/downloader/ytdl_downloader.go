package downloader

import (
	"path/filepath"
	"time"

	"github.com/BrianAllred/goydl"
)

// YoutubeDLDownloader is youtube-dl downloader
type YoutubeDLDownloader struct {
	Downloader
	URL            string
	Destination    string
	ReportInterval time.Duration
	IsDone         chan bool
	LastError      error
	progressChan   chan DownloadProgress
	BinaryPath     string
	Label          string
}

// YoutubeDLDownloaderParams params for setting up a youtube-dl Downloader
type YoutubeDLDownloaderParams struct {
	URL            string
	Destination    string
	ReportInterval time.Duration
	BinaryPath     string
}

// NewYoutubeDLDownloader creates a youtube-dl downloader
func NewYoutubeDLDownloader(params YoutubeDLDownloaderParams) Downloader {
	return &YoutubeDLDownloader{
		URL:            params.URL,
		Destination:    params.Destination,
		ReportInterval: params.ReportInterval,
		BinaryPath:     params.BinaryPath,
		IsDone:         make(chan bool),
		progressChan:   make(chan DownloadProgress),
	}
}

// PreProcess pre process a url, e.g. getting if this url support resume download
func (u YoutubeDLDownloader) PreProcess() {
}

// Process will start a download process
func (u YoutubeDLDownloader) Process() {
	ytdl := goydl.NewYoutubeDl()
	ytdl.YoutubeDlPath = u.BinaryPath
	ytdl.Options.Output.Value = filepath.Join(u.Destination, "%(title)s.%(ext)s")

	cmd, err := ytdl.Download(u.URL)

	if err != nil {
		u.LastError = err
		return
	}

	progress := DownloadProgress{
		Label:    u.Label,
		Progress: 0.0,
	}
	u.progressChan <- progress
	cmd.Wait()

}

// PostProcess will clean up files
func (u YoutubeDLDownloader) PostProcess() {
	// Remove meta data
}

// Report will return progress channel
func (u YoutubeDLDownloader) Report() chan DownloadProgress {
	return u.progressChan
}

// Done specify this task is done
func (u YoutubeDLDownloader) Done() chan bool {
	return u.IsDone
}

// Error specify what error occur in this task
func (u YoutubeDLDownloader) Error() error {
	return u.LastError
}

// Identifier describe downloader identity
func (u YoutubeDLDownloader) Identifier() string {
	return "yotube-dl"
}