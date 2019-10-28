package downloader

// URLDownloader download file from url
type URLDownloader struct {
	Downloader
}

// PreProcess pre process a url, e.g. getting if this url support resume download
func (u URLDownloader) PreProcess() {

}

// Process will start a download process
func (u URLDownloader) Process() {

}

// PostProcess will clean up files
func (u URLDownloader) PostProcess() {

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
