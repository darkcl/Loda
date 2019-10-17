package models

// Download stores data about a download task
type Download struct {
	Status         int
	TotalLength    int64
	BytesCompleted int64
	BytesUpload    int64
	DownloadSpeed  int
	UploadSpeed    int
	NumPieces      int
	Connections    int
	BitField       string
	InfoHash       string
	MetaInfo       MetaInfo
	Files          []File
}
