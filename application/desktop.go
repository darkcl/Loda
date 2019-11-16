package application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/darkcl/loda/application/repositories"

	"github.com/asdine/storm/v3"
	"github.com/darkcl/loda/application/controllers"
	"github.com/mitchellh/go-homedir"

	"github.com/darkcl/loda/lib/ipc"
	"github.com/darkcl/webview"
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

	Controllers []controllers.Controller

	db *storm.DB
}

// WillLaunch call before application is launch
func (d *DesktopApplication) WillLaunch(mode string, configuration map[string]string) {
	d.BaseApplication.WillLaunch(mode, configuration)

	d.ApplicationName = "Loda"
	d.IPCMain = &ipc.Main{
		Callback: map[string]ipc.EventCallback{},
	}

	dir, err := ioutil.TempDir("", d.ApplicationName)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	d.AssetsDir = dir
	d.db = d.createDB()

	d.embeddedYTDL()

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

	d.Controllers = []controllers.Controller{
		&controllers.LinkController{},
		&controllers.DownloadController{
			Repository: repositories.NewDownloadRepository(d.db),
		},
	}
}

// DidFinishLaunching call after all application launch logic is completed
func (d *DesktopApplication) DidFinishLaunching() {
	d.BaseApplication.DidFinishLaunching()

	launchContext := make(map[string]interface{})
	launchContext["ipc"] = d.IPCMain

	for _, con := range d.Controllers {
		con.Load(launchContext)
	}

	d.Window.Run()
}

// WillTerminate call before application Terminate
func (d *DesktopApplication) WillTerminate() {
	d.BaseApplication.WillTerminate()
	os.RemoveAll(d.AssetsDir)
	d.db.Close()
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

func workspaceDir() string {
	path, err := homedir.Dir()

	if err != nil {
		panic(err)
	}
	defaultWorkspace := filepath.Join(path, ".loda")
	if _, err := os.Stat(defaultWorkspace); os.IsNotExist(err) {
		os.Mkdir(defaultWorkspace, os.ModePerm)
	}

	return defaultWorkspace
}

func (d *DesktopApplication) createDB() *storm.DB {
	defaultWorkspace := workspaceDir()

	dbPath := filepath.Join(defaultWorkspace, "data.db")

	db, err := storm.Open(dbPath)

	if err != nil {
		panic(err)
	}

	return db
}

func (d *DesktopApplication) embeddedYTDL() {
	defaultWorkspace := workspaceDir()
	binName := "youtube-dl"

	ytdl := mewn.String("./bin/youtube-dl")
	if runtime.GOOS == "windows" {
		ytdl = mewn.String("./bin/youtube-dl.exe")
		binName = "youtube-dl.exe"
	}

	ytdlPath := filepath.Join(defaultWorkspace, binName)

	if err := ioutil.WriteFile(ytdlPath, []byte(ytdl), 0777); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
