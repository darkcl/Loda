package application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/darkcl/loda/helpers"
	"github.com/darkcl/loda/ipc"
	"github.com/darkcl/loda/lib/webview"
	"github.com/leaanthony/mewn"
)

// DesktopApplication implements desktop application life cycle
type DesktopApplication struct {
	BaseApplication

	ApplicationName string

	LaunchURL string
	AssetsDir string

	IPCMain *ipc.Main
	Window  webview.WebView
}

// WillLaunch call before application is launch
func (d *DesktopApplication) WillLaunch(mode string, configuration map[string]string) {
	d.BaseApplication.WillLaunch(mode, configuration)

	d.ApplicationName = "Loda"

	dir, err := ioutil.TempDir("", d.ApplicationName)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	d.AssetsDir = dir

	switch d.Mode {
	case "release":
		d.LaunchURL = d.bundleAssets()
		d.Window = d.createWindow(false)
		d.IPCMain.SetView(d.Window)
	case "debug":
		d.LaunchURL = fmt.Sprintf("http://%s:%s", configuration["host"], configuration["port"])
		d.Window = d.createWindow(true)
		d.IPCMain.SetView(d.Window)
	}
}

// DidFinishLaunching call after all application launch logic is completed
func (d *DesktopApplication) DidFinishLaunching() {
	d.BaseApplication.DidFinishLaunching()

	d.Window.Run()

	d.IPCMain.On(
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

// WillTerminate call before application Terminate
func (d *DesktopApplication) WillTerminate() {
	d.BaseApplication.WillTerminate()
	os.RemoveAll(d.AssetsDir)
}

func (d *DesktopApplication) bundleAssets() string {
	js := mewn.String("./ui/dist/bundle.min.js")
	indexHTML := mewn.String("./ui/dist/index.html")

	tmpIndex := filepath.Join(d.AssetsDir, "index.html")
	if err := ioutil.WriteFile(tmpIndex, []byte(indexHTML), 0666); err != nil {
		log.Fatal(err)
		panic(err)
	}

	abs, err := filepath.Abs(tmpIndex)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	tmpJs := filepath.Join(d.AssetsDir, "bundle.min.js")
	if err := ioutil.WriteFile(tmpJs, []byte(js), 0666); err != nil {
		log.Fatal(err)
		panic(err)
	}

	return "file://" + abs
}

func (d *DesktopApplication) createWindow(debug bool) webview.WebView {
	return webview.New(webview.Settings{
		Title:                  "Loda",
		URL:                    d.LaunchURL,
		Resizable:              true,
		Width:                  1024,
		Height:                 768,
		ExternalInvokeCallback: d.handleRPC,
		Debug:                  debug,
	})
}

func (d *DesktopApplication) handleRPC(w webview.WebView, data string) {
	var message ipc.Message
	err := json.Unmarshal([]byte(data), &message)

	if err != nil {
		fmt.Printf("Error on handle rpc data: %v\n", err)
		return
	}

	d.IPCMain.Trigger(message)
}
