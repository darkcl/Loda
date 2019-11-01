package downloader

import "time"

// Downloader interface describe a downloader
type Downloader interface {
	PreProcess()
	Process()
	PostProcess()

	Report() chan DownloadProgress
	Done() chan bool
	Error() error
	Identifier() string
}

// ProgressCallback is the callback to report download progress
type ProgressCallback = func(progress *DownloadProgress)

// DownloadProgress describe current download progress
type DownloadProgress struct {
	Label          string    `json:"label"`
	ETA            time.Time `json:"eta"`
	StartAt        time.Time `json:"startAt"`
	EndAt          time.Time `json:"endAt"`
	BytesComplete  int64     `json:"bytesComplete"`
	BytesPerSecond float64   `json:"bytesPerSecond"`
	Progress       float64   `json:"progress"`
}
