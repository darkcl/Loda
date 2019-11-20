package downloader

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/BrianAllred/goydl"
)

// YoutubeDLDownloader is youtube-dl downloader
type YoutubeDLDownloader struct {
	Downloader
	URL            string
	Destination    string
	ReportInterval time.Duration
	IsDone         chan bool
	LastError      error
	progressChan   chan DownloadProgress
	BinaryPath     string
	Label          string
	prevProgress   DownloadProgress
}

// YoutubeDLDownloaderParams params for setting up a youtube-dl Downloader
type YoutubeDLDownloaderParams struct {
	URL            string
	Destination    string
	ReportInterval time.Duration
	BinaryPath     string
}

type outputProcessor struct {
	io.Writer
	Loader *YoutubeDLDownloader
}

func (o *outputProcessor) Write(p []byte) (n int, err error) {
	go func(input string) {
		progress := o.Loader.ParseProgress(input)
		o.Loader.ReportProgress(progress)
	}(string(p))
	return len(p), nil
}

// NewYoutubeDLDownloader creates a youtube-dl downloader
func NewYoutubeDLDownloader(params YoutubeDLDownloaderParams) Downloader {
	return &YoutubeDLDownloader{
		URL:            params.URL,
		Destination:    params.Destination,
		ReportInterval: params.ReportInterval,
		BinaryPath:     params.BinaryPath,
		IsDone:         make(chan bool),
		progressChan:   make(chan DownloadProgress),
	}
}

// PreProcess pre process a url, e.g. getting if this url support resume download
func (u YoutubeDLDownloader) PreProcess() {
}

// Process will start a download process
func (u *YoutubeDLDownloader) Process() {
	ytdl := goydl.NewYoutubeDl()
	ytdl.YoutubeDlPath = u.BinaryPath
	ytdl.Options.Output.Value = filepath.Join(u.Destination, "%(title)s.%(ext)s")

	cmd, err := ytdl.Download(u.URL)

	if err != nil {
		u.LastError = err
		return
	}

	go io.Copy(&outputProcessor{
		Loader: u,
	}, ytdl.Stdout)
	go io.Copy(os.Stderr, ytdl.Stderr)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("Download complete")
	}

	u.IsDone <- true
	u.OnComplete()
}

// OnComplete will call on download task is completed
func (u YoutubeDLDownloader) OnComplete() {
	fmt.Println("YoutubeDLDownloader OnComplete")
}

// PostProcess will clean up files
func (u YoutubeDLDownloader) PostProcess() {
	// Remove meta data
}

// Report will return progress channel
func (u *YoutubeDLDownloader) Report() chan DownloadProgress {
	return u.progressChan
}

// Done specify this task is done
func (u *YoutubeDLDownloader) Done() chan bool {
	return u.IsDone
}

// Error specify what error occur in this task
func (u YoutubeDLDownloader) Error() error {
	return u.LastError
}

// Identifier describe downloader identity
func (u YoutubeDLDownloader) Identifier() string {
	return "youtube-dl"
}

// ParseProgress parse youtube-dl output to progress object
func (u *YoutubeDLDownloader) ParseProgress(input string) DownloadProgress {
	result := DownloadProgress{}
	progressRegex := regexp.MustCompile(`[0-9.]+\%`)
	progress, err := u.parseFloat(progressRegex.FindString(input), 64)

	if err != nil {
		// Use previous progress
		return u.prevProgress
	}
	result.Progress = progress
	u.prevProgress = result
	return result
}

// ReportProgress is called when output processor is parsed output
func (u *YoutubeDLDownloader) ReportProgress(progress DownloadProgress) {
	u.progressChan <- progress
}

func (u *YoutubeDLDownloader) parseFloat(input string, bitSize int) (f float64, err error) {
	i := strings.Index(input, "%")
	if i < 0 {
		return 0, fmt.Errorf("ParseFloatPercent: percentage sign not found")
	}
	f, err = strconv.ParseFloat(input[:i], bitSize)
	if err != nil {
		return 0, err
	}
	return f / 100.0, nil
}
