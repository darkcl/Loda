package controllers

import (
	"github.com/darkcl/webview"

	"github.com/darkcl/loda/lib/ipc"
)

// FolderController implements all folder related logic (open, select...)
type FolderController struct {
	Controller
	ipcMain *ipc.Main
}

// Load is called when application is loaded
func (f *FolderController) Load(context map[string]interface{}) {
	// Load IPC
	ipcMain, ok := context["ipc"].(*ipc.Main)
	f.ipcMain = ipcMain
	if ok == false {
		panic("Require IPC Processor")
	}

	ipcMain.On(
		"request.open_directory",
		f.onOpenDirectory)
}

func (f FolderController) onOpenDirectory(event string, value interface{}) interface{} {
	w := f.ipcMain.CurrentView()
	result := w.Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, "Open directory", "")
	if result != "" {
		f.ipcMain.Send("response.open_directory", map[string]string{
			"directory": result,
		})
		return nil
	}
	return nil
}
