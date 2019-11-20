package downloader

import (
	"fmt"
	"time"

	"github.com/anacrolix/torrent"
)

// TorrentDownloader downloads torrent and magnet link
type TorrentDownloader struct {
	Downloader
	TorrentFile    string
	Destination    string
	ReportInterval time.Duration
	IsDone         chan bool
	LastError      error
	progressChan   chan DownloadProgress
	Label          string
}

// TorrentDownloaderParams params for torrent downloader
type TorrentDownloaderParams struct {
	TorrentFile    string
	Destination    string
	ReportInterval time.Duration
	Label          string
}

// NewTorrentDownloader creates a torrent downloader
func NewTorrentDownloader(params TorrentDownloaderParams) Downloader {
	return &TorrentDownloader{
		TorrentFile:    params.TorrentFile,
		Destination:    params.Destination,
		ReportInterval: params.ReportInterval,
		Label:          params.Label,
		IsDone:         make(chan bool),
		LastError:      nil,
		progressChan:   make(chan DownloadProgress),
	}
}

// PreProcess pre process a url, e.g. getting if this url support resume download
func (u TorrentDownloader) PreProcess() {
}

// Process will start a download process
func (u TorrentDownloader) Process() {
	defer func() {
		if r := recover(); r != nil {
			u.LastError = r.(error)
			u.IsDone <- true
			fmt.Printf("[TorrentDownloader] %v\n", r)
		}
	}()

	clientConfig := torrent.NewDefaultClientConfig()
	clientConfig.DataDir = u.Destination
	client, err := torrent.NewClient(clientConfig)

	if err != nil {
		panic(err)
	}

	t, err := client.AddTorrentFromFile(u.TorrentFile)

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
		for {
			<-t.GotInfo()
			select {
			case <-ticker.C:
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
	}()

	go func() {
		<-t.GotInfo()
		t.DownloadAll()
	}()
}

// OnComplete will call on download task is completed
func (u TorrentDownloader) OnComplete() {
}

// PostProcess will clean up files
func (u TorrentDownloader) PostProcess() {
	// Remove meta data
}

// Report will return progress channel
func (u TorrentDownloader) Report() chan DownloadProgress {
	return u.progressChan
}

// Done specify this task is done
func (u TorrentDownloader) Done() chan bool {
	return u.IsDone
}

// Error specify what error occur in this task
func (u TorrentDownloader) Error() error {
	return u.LastError
}

// Identifier describe downloader identity
func (u TorrentDownloader) Identifier() string {
	return "torrent"
}
