package models

// MetaInfo stores meta information about a Bittorrent download task
type MetaInfo struct {
	Name         string
	AnnounceList []string
	Comment      string
	CreationUnix int64
	Mode         string
}
