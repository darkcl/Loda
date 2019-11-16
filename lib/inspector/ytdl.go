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
	ytdl.Options.DumpJSON.Value = true
	cmd, err := ytdl.Download(input)

	if err != nil {
		return nil, err
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	return ytdl.Info, nil
}
