package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/darkcl/loda/application/services"

	"github.com/darkcl/loda/application/repositories"

	"github.com/darkcl/loda/lib/downloader"
	"github.com/segmentio/ksuid"

	"github.com/darkcl/loda/lib/ipc"
	"github.com/darkcl/loda/lib/matcher"
)

// DownloadController implements all download related logic
type DownloadController struct {
	Controller

	Matchers   []matcher.Matcher
	Repository repositories.DownloadRepository

	progressService services.DownloadProgressService
}

// DownloadRequest - Download request model
type DownloadRequest struct {
	URL         string `json:"url"`
	Destination string `json:"destination"`
}

// DownloadProgressRequest - Download progress request model
type DownloadProgressRequest struct {
	ID int `json:"id"`
}

// Load is called when application is loaded
func (d *DownloadController) Load(context map[string]interface{}) {
	// Load Matcher
	d.Matchers = []matcher.Matcher{
		&matcher.URLMatcher{},
	}

	d.progressService = services.NewDownloadProgessService(d.Repository)

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

	ipcMain.On(
		"request.download_progress",
		func(event string, value interface{}) interface{} {
			payload, ok := value.(string)

			if ok == false {
				ipcMain.Send("error.download_progress", map[string]string{
					"error": "Value is not a string",
				})
			}

			var request DownloadProgressRequest
			err := json.Unmarshal([]byte(payload), &request)
			if err != nil {
				ipcMain.Send("error.download_progress", map[string]string{
					"error": err.Error(),
				})
			}

			task, err := d.Repository.FindOne(request.ID)

			if err != nil {
				ipcMain.Send("error.download_progress", map[string]string{
					"error": "Progress not found",
				})
				return nil
			}

			ipcMain.Send("progress.download", task.Progress)
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
		label := ksuid.New().String()
		ipcMain.Send("download_label", label)
		go d.startURLDownloader(request, ipcMain, label)
	default:
		ipcMain.Send("error.create_download", map[string]string{
			"error": fmt.Sprintf("No downloader of this match: %s", downloaderID),
		})
	}
}

func (d DownloadController) startURLDownloader(request DownloadRequest, ipcMain *ipc.Main, label string) {
	interval := 1000 * time.Millisecond

	task, err := d.Repository.Create(request.Destination, "url-download")

	if err != nil {
		ipcMain.Send("error.create_download", map[string]string{
			"error": err.Error(),
		})
		return
	}

	loader := downloader.NewURLDownloader(downloader.URLDownloaderParams{
		URL:              request.URL,
		Label:            label,
		Destination:      request.Destination,
		NumOfConnections: 1,
		IsResumable:      true,
		ReportInterval:   interval,
	})

	defer loader.PostProcess()

	loader.PreProcess()
	go loader.Process()

	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-loader.Done():
			return
		case <-t.C:
			p := <-loader.Report()
			d.progressService.UpdateProgress(task, p)
			fmt.Printf("Progress: %s\n", p.Label)
		}
	}
}
