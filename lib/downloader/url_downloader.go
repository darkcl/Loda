package downloader

// URLDownloader download file from url
type URLDownloader struct {
	Downloader
	URL              string
	Destination      string
	NumOfConnections int
	IsResumable      bool
}

// URLDownloaderParams params for setting up a URL Downloader
type URLDownloaderParams struct {
	URL              string
	Destination      string
	NumOfConnections int
	IsResumable      bool
}

// NewURLDownloader creates an url downloader
func NewURLDownloader(parms URLDownloaderParams) *URLDownloader {
	return &URLDownloader{
		URL:              parms.URL,
		Destination:      parms.Destination,
		NumOfConnections: parms.NumOfConnections,
		IsResumable:      parms.IsResumable,
	}
}

// PreProcess pre process a url, e.g. getting if this url support resume download
func (u URLDownloader) PreProcess() {
	// Get URL headers

	// Get Destination

	// Get Resumable Meta data
}

// Process will start a download process
func (u URLDownloader) Process() {

}

// PostProcess will clean up files
func (u URLDownloader) PostProcess() {
	// Remove meta data
}

// Done specify this task is done
func (u URLDownloader) Done() bool {
	return false
}

// Error specify what error occur in this task
func (u URLDownloader) Error() error {
	return nil
}

// Identifier describe downloader identity
func (u URLDownloader) Identifier() string {
	return "url"
}
