package downloader

// Downloader interface describe a downloader
type Downloader interface {
	PreProcess()
	Process()
	PostProcess()

	Done() bool
	Error() error
	Identifier() string
}
