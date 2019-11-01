package controllers

import (
	"fmt"

	"github.com/darkcl/loda/helpers"
	"github.com/darkcl/loda/lib/ipc"
)

// LinkController implements all links related logic
type LinkController struct {
	Controller
}

// Load is called when application is loaded
func (l *LinkController) Load(context map[string]interface{}) {
	ipcMain, ok := context["ipc"].(*ipc.Main)

	if ok == false {
		panic("Require IPC Processor")
	}

	ipcMain.On(
		"openlink",
		func(event string, value interface{}) interface{} {
			if value == nil {
				fmt.Printf("[openlink] value not provided\n")
				return nil
			}

			fmt.Printf("Open Link: %s", value.(string))
			url := value.(string)
			helpers.OpenBrowser(url)
			return nil
		})
}
