package downloader

import (
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/segmentio/ksuid"
)

// URLDownloader download file from url
type URLDownloader struct {
	Downloader
	URL              string
	Destination      string
	NumOfConnections int
	IsResumable      bool
	ReportInterval   time.Duration
	IsDone           bool
	LastError        error
}

// URLDownloaderParams params for setting up a URL Downloader
type URLDownloaderParams struct {
	URL              string
	Destination      string
	NumOfConnections int
	IsResumable      bool
	ReportInterval   time.Duration
}

// NewURLDownloader creates an url downloader
func NewURLDownloader(params URLDownloaderParams) *URLDownloader {
	return &URLDownloader{
		URL:              params.URL,
		Destination:      params.Destination,
		NumOfConnections: params.NumOfConnections,
		IsResumable:      params.IsResumable,
		ReportInterval:   params.ReportInterval,
		IsDone:           false,
		LastError:        nil,
	}
}

// PreProcess pre process a url, e.g. getting if this url support resume download
func (u URLDownloader) PreProcess() {
}

// Process will start a download process
func (u URLDownloader) Process(report ProgressCallback) {
	label := ksuid.New().String()
	client := grab.NewClient()
	req, _ := grab.NewRequest(u.Destination, u.URL)
	resp := client.Do(req)
	t := time.NewTicker(u.ReportInterval)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			progress := DownloadProgress{
				Label:          label,
				ETA:            resp.ETA(),
				StartAt:        resp.Start,
				EndAt:          resp.End,
				BytesComplete:  resp.BytesComplete(),
				BytesPerSecond: resp.BytesPerSecond(),
				Progress:       resp.Progress(),
			}
			report(progress)

		case <-resp.Done:
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		u.LastError = err
	}

	u.IsDone = true
}

// PostProcess will clean up files
func (u URLDownloader) PostProcess() {
	// Remove meta data
}

// Done specify this task is done
func (u URLDownloader) Done() bool {
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
