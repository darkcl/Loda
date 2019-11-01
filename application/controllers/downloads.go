package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/darkcl/loda/lib/downloader"

	"github.com/darkcl/loda/lib/ipc"
	"github.com/darkcl/loda/lib/matcher"
)

// DownloadController implements all download related logic
type DownloadController struct {
	Controller

	Matchers []matcher.Matcher
}

// DownloadRequest - Download request model
type DownloadRequest struct {
	URL         string `json:"url"`
	Destination string `json:"destination"`
}

// Load is called when application is loaded
func (d *DownloadController) Load(context map[string]interface{}) {
	// Load Matcher
	d.Matchers = []matcher.Matcher{
		&matcher.URLMatcher{},
	}

	// Load IPC
	ipcMain, ok := context["ipc"].(*ipc.Main)

	if ok == false {
		panic("Require IPC Processor")
	}

	ipcMain.On(
		"request.create_download",
		func(event string, value interface{}) interface{} {
			payload, ok := value.(string)

			if ok == false {
				ipcMain.Send("error.create_download", map[string]string{
					"error": "Value is not a string",
				})
			}

			var request DownloadRequest
			err := json.Unmarshal([]byte(payload), &request)
			if err != nil {
				ipcMain.Send("error.create_download", map[string]string{
					"error": err.Error(),
				})
			}

			d.CreateDownloadTask(request, ipcMain)
			return nil
		})
}

// CreateDownloadTask will create a download taks and start notify download progress
func (d *DownloadController) CreateDownloadTask(request DownloadRequest, ipcMain *ipc.Main) {
	// Match The URL
	downloaderID := ""

	for _, m := range d.Matchers {
		matched, _ := m.Process(request.URL)
		if matched == true {
			downloaderID = m.Identifier()
		}
	}

	if downloaderID == "" {
		ipcMain.Send("error.create_download", map[string]string{
			"error": "No matching downloader",
		})
	}

	// Create Downloader
	switch downloaderID {
	case "url":
		d.startURLDownloader(request, ipcMain)
	default:
		ipcMain.Send("error.create_download", map[string]string{
			"error": fmt.Sprintf("No downloader of this match: %s", downloaderID),
		})
	}
}

func (d DownloadController) startURLDownloader(request DownloadRequest, ipcMain *ipc.Main) {
	loader := downloader.NewURLDownloader(downloader.URLDownloaderParams{
		URL:              request.URL,
		Destination:      request.Destination,
		NumOfConnections: 1,
		IsResumable:      true,
		ReportInterval:   500 * time.Millisecond,
	})

	defer loader.PostProcess()

	loader.PreProcess()
	loader.Process(func(progress downloader.DownloadProgress) {
		ipcMain.Send("progress.download", progress)
	})
}
