package models

// File stores information about a downloadable file
type File struct {
	Index           int
	Name            string
	Length          int64
	CompletedLength int64
	Selected        bool
}
