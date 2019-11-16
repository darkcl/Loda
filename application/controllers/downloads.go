package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/darkcl/loda/application/models"
	"github.com/darkcl/loda/application/services"

	"github.com/darkcl/loda/application/repositories"

	"github.com/darkcl/loda/lib/downloader"

	"github.com/darkcl/loda/lib/ipc"
)

// DownloadController implements all download related logic
type DownloadController struct {
	Controller

	Repository repositories.DownloadRepository

	progressService services.DownloadProgressService
	matcherService  services.MatcherService
	ipcMain         *ipc.Main
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
	// Load Service
	d.progressService = services.NewDownloadProgessService(d.Repository)
	d.matcherService = services.NewMatcherService()

	// Load IPC
	ipcMain, ok := context["ipc"].(*ipc.Main)
	d.ipcMain = ipcMain

	if ok == false {
		panic("Require IPC Processor")
	}

	ipcMain.On(
		"request.create_download",
		d.onCreateDownload)

	ipcMain.On(
		"request.download_progress",
		d.onDownloadProgress)

	ipcMain.On(
		"request.download_list",
		d.onDownloadList)
}

func (d DownloadController) onDownloadList(event string, value interface{}) interface{} {
	tasks, err := d.Repository.All()
	if err != nil {
		d.ipcMain.Send("error.download_list", map[string]string{
			"error": err.Error(),
		})
		return nil
	}
	d.ipcMain.Send("response.download_list", tasks)
	return nil
}

func (d DownloadController) onDownloadProgress(event string, value interface{}) interface{} {
	payload, ok := value.(string)

	if ok == false {
		d.ipcMain.Send("error.download_progress", map[string]string{
			"error": "Value is not a string",
		})
		return nil
	}

	var request DownloadProgressRequest
	err := json.Unmarshal([]byte(payload), &request)
	if err != nil {
		d.ipcMain.Send("error.download_progress", map[string]string{
			"error": err.Error(),
		})
		return nil
	}

	task, err := d.Repository.FindOne(request.ID)

	if err != nil {
		d.ipcMain.Send("error.download_progress", map[string]string{
			"error": "Progress not found",
		})
		return nil
	}

	d.ipcMain.Send("response.progress.download", task.Progress)

	if task.IsDone == true {
		fmt.Printf("Task is done\n")
		d.ipcMain.Send("response.progress.download.done", task)
	}
	return nil
}

func (d DownloadController) onCreateDownload(event string, value interface{}) interface{} {
	payload, ok := value.(string)

	if ok == false {
		d.ipcMain.Send("error.create_download", map[string]string{
			"error": "Value is not a string",
		})
		return nil
	}

	var request DownloadRequest
	err := json.Unmarshal([]byte(payload), &request)
	if err != nil {
		d.ipcMain.Send("error.create_download", map[string]string{
			"error": err.Error(),
		})
		return nil
	}

	d.CreateDownloadTask(request, d.ipcMain)
	return nil
}

// CreateDownloadTask will create a download taks and start notify download progress
func (d *DownloadController) CreateDownloadTask(request DownloadRequest, ipcMain *ipc.Main) {
	// Match The URL
	downloader, err := d.matcherService.Match(request.URL, request.Destination)

	if err != nil {
		ipcMain.Send("error.create_download", map[string]string{
			"error": err.Error(),
		})
		return
	}

	task, err := d.Repository.Create(request.Destination, downloader.Identifier())

	if err != nil {
		ipcMain.Send("error.create_download", map[string]string{
			"error": err.Error(),
		})
		return
	}

	ipcMain.Send("download_label", task.ID)
	go d.startDownloader(downloader, ipcMain, task)
}

func (d DownloadController) startDownloader(loader downloader.Downloader, ipcMain *ipc.Main, task *models.DownloadTask) {
	interval := 1000 * time.Millisecond
	defer loader.PostProcess()

	loader.PreProcess()
	go loader.Process()

	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-loader.Done():
			d.progressService.MarkDone(task)
			return
		case <-t.C:
			p := <-loader.Report()
			d.progressService.UpdateProgress(task, p)
			fmt.Printf("Progress: %s\n", p.Label)
		}
	}
}
