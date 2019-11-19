package downloader

import "time"

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
