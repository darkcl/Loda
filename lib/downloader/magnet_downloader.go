package downloader

import (
	"fmt"
	"time"

	"github.com/anacrolix/torrent"
)

// MagnetDownloader downloads torrent and magnet link
type MagnetDownloader struct {
	Downloader
	MagnetURI      string
	Destination    string
	ReportInterval time.Duration
	IsDone         chan bool
	LastError      error
	progressChan   chan DownloadProgress
	Label          string
}

// MagnetDownloaderParams params for torrent downloader
type MagnetDownloaderParams struct {
	MagnetURI      string
	Destination    string
	ReportInterval time.Duration
	Label          string
}

// NewMagnetDownloader creates a torrent downloader
func NewMagnetDownloader(params MagnetDownloaderParams) Downloader {
	return &MagnetDownloader{
		MagnetURI:      params.MagnetURI,
		Destination:    params.Destination,
		ReportInterval: params.ReportInterval,
		Label:          params.Label,
		IsDone:         make(chan bool),
		LastError:      nil,
		progressChan:   make(chan DownloadProgress),
	}
}

// PreProcess pre process a url, e.g. getting if this url support resume download
func (u MagnetDownloader) PreProcess() {
}

// Process will start a download process
func (u MagnetDownloader) Process() {
	defer func() {
		if r := recover(); r != nil {
			u.LastError = r.(error)
			u.IsDone <- true
			fmt.Printf("[MagnetDownloader] %v\n", r)
		}
	}()

	clientConfig := torrent.NewDefaultClientConfig()
	clientConfig.DataDir = u.Destination
	client, err := torrent.NewClient(clientConfig)

	if err != nil {
		panic(err)
	}

	t, err := client.AddMagnet(u.MagnetURI)

	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(u.ReportInterval)
	defer ticker.Stop()

	progress := DownloadProgress{
		Label:    u.Label,
		Progress: 0.0,
	}

	go func() {
		<-t.GotInfo()
		t.DownloadAll()
	}()

	for {
		select {
		case <-ticker.C:
			<-t.GotInfo()
			if t.BytesCompleted() == t.Info().TotalLength() {
				u.OnComplete()
				u.IsDone <- true
				return
			}

			progress.BytesComplete = t.BytesCompleted()
			progress.Progress = float64(t.BytesCompleted()) / float64(t.Info().TotalLength())
			u.progressChan <- progress
		}
	}
}

// OnComplete will call on download task is completed
func (u MagnetDownloader) OnComplete() {
}

// PostProcess will clean up files
func (u MagnetDownloader) PostProcess() {
	// Remove meta data
}

// Report will return progress channel
func (u MagnetDownloader) Report() chan DownloadProgress {
	return u.progressChan
}

// Done specify this task is done
func (u MagnetDownloader) Done() chan bool {
	return u.IsDone
}

// Error specify what error occur in this task
func (u MagnetDownloader) Error() error {
	return u.LastError
}

// Identifier describe downloader identity
func (u MagnetDownloader) Identifier() string {
	return "magnet"
}
