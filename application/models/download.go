package models

import "time"

// DownloadProgress describe a download progress
type DownloadProgress struct {
	Label          string    `json:"label"`
	ETA            time.Time `json:"eta"`
	StartAt        time.Time `json:"startAt"`
	EndAt          time.Time `json:"endAt"`
	BytesComplete  int64     `json:"bytesComplete"`
	BytesPerSecond float64   `json:"bytesPerSecond"`
	Progress       float64   `json:"progress"`
}

// DownloadTask describe a download task
type DownloadTask struct {
	ID          int              `storm:"id,increment"`
	Destination string           `storm:"index"`
	TaskType    string           `storm:"index"`
	IsDone      bool             `storm:"index"`
	Progress    DownloadProgress `storm:"inline"`
}
