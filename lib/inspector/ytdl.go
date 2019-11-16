package inspector

import "github.com/BrianAllred/goydl"

// YoutubeDLInspector will inspect on a file or a url for meta data
type YoutubeDLInspector struct {
	Inspector
	binaryPath string
}

// NewYoutubeDLInspector create youtube-dl inspector
func NewYoutubeDLInspector(binaryPath string) Inspector {
	return &YoutubeDLInspector{
		binaryPath: binaryPath,
	}
}

// Process will process input file path / url and return meta data
func (y YoutubeDLInspector) Process(input string) (interface{}, error) {
	ytdl := goydl.NewYoutubeDl()
	ytdl.YoutubeDlPath = y.binaryPath
	_, err := ytdl.Download(input)
	return ytdl.Info, err
}
