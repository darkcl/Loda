package downloader

import (
	"time"

	"github.com/cavaliercoder/grab"
)

// URLDownloader download file from url
type URLDownloader struct {
	Downloader
	URL              string
	Destination      string
	NumOfConnections int
	IsResumable      bool
	ReportInterval   time.Duration
	IsDone           chan bool
	LastError        error
	progressChan     chan DownloadProgress
	Label            string
}

// URLDownloaderParams params for setting up a URL Downloader
type URLDownloaderParams struct {
	URL              string
	Destination      string
	NumOfConnections int
	IsResumable      bool
	ReportInterval   time.Duration
	Label            string
}

// NewURLDownloader creates an url downloader
func NewURLDownloader(params URLDownloaderParams) Downloader {
	return &URLDownloader{
		URL:              params.URL,
		Destination:      params.Destination,
		NumOfConnections: params.NumOfConnections,
		IsResumable:      params.IsResumable,
		ReportInterval:   params.ReportInterval,
		Label:            params.Label,
		IsDone:           make(chan bool),
		LastError:        nil,
		progressChan:     make(chan DownloadProgress),
	}
}

// PreProcess pre process a url, e.g. getting if this url support resume download
func (u URLDownloader) PreProcess() {
}

// Process will start a download process
func (u URLDownloader) Process() {
	client := grab.NewClient()
	req, _ := grab.NewRequest(u.Destination, u.URL)
	resp := client.Do(req)

	t := time.NewTicker(u.ReportInterval)
	defer t.Stop()

	progress := DownloadProgress{
		Label:          u.Label,
		ETA:            resp.ETA(),
		StartAt:        resp.Start,
		EndAt:          resp.End,
		BytesComplete:  resp.BytesComplete(),
		BytesPerSecond: resp.BytesPerSecond(),
		Progress:       resp.Progress(),
	}

	for {
		select {
		case <-resp.Done:
			// check for errors
			if err := resp.Err(); err != nil {
				u.LastError = err
			}
			u.OnComplete()
			u.IsDone <- true
			return
		case <-t.C:
			progress.ETA = resp.ETA()
			progress.StartAt = resp.Start
			progress.EndAt = resp.End
			progress.BytesComplete = resp.BytesComplete()
			progress.BytesPerSecond = resp.BytesPerSecond()
			progress.Progress = resp.Progress()
			u.progressChan <- progress
		}
	}
}

// OnComplete will call on download task is completed
func (u URLDownloader) OnComplete() {
}

// PostProcess will clean up files
func (u URLDownloader) PostProcess() {
	// Remove meta data
}

// Report will return progress channel
func (u URLDownloader) Report() chan DownloadProgress {
	return u.progressChan
}

// Done specify this task is done
func (u URLDownloader) Done() chan bool {
	return u.IsDone
}

// Error specify what error occur in this task
func (u URLDownloader) Error() error {
	return u.LastError
}

// Identifier describe downloader identity
func (u URLDownloader) Identifier() string {
	return "url"
}
