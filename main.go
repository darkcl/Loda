package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/darkcl/loda/helpers"
	"github.com/darkcl/loda/ipc"

	webview "github.com/darkcl/webview"
	"github.com/leaanthony/mewn"
)

func handleRPC(w webview.WebView, data string) {
	var message ipc.Message
	err := json.Unmarshal([]byte(data), &message)

	if err != nil {
		fmt.Printf("Error on handle rpc data: %v\n", err)
		return
	}

	ipcMain := ipc.SharedMain()
	ipcMain.Trigger(message)
}

func bundleAssets() string {
	js := mewn.String("./ui/dist/bundle.min.js")
	indexHTML := mewn.String("./ui/dist/index.html")

	dir, err := ioutil.TempDir("", "skeleton")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer os.RemoveAll(dir)
	tmpIndex := filepath.Join(dir, "index.html")
	if err := ioutil.WriteFile(tmpIndex, []byte(indexHTML), 0666); err != nil {
		log.Fatal(err)
		panic(err)
	}
	abs, err := filepath.Abs(tmpIndex)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	tmpJs := filepath.Join(dir, "bundle.min.js")
	if err := ioutil.WriteFile(tmpJs, []byte(js), 0666); err != nil {
		log.Fatal(err)
		panic(err)
	}

	return abs
}

func createWindow(mode string, host string, port int) webview.WebView {
	switch mode {
	case "release":
		fmt.Println("Release Mode")
		abs := bundleAssets()

		return webview.New(webview.Settings{
			Title:                  "Loda",
			URL:                    "file://" + abs,
			Resizable:              true,
			Width:                  1024,
			Height:                 768,
			ExternalInvokeCallback: handleRPC,
			Debug:                  true,
		})
	case "debug":
		fmt.Println("Debug Mode")
		webpackUrl := fmt.Sprintf("http://%s:%v", host, port)
		return webview.New(webview.Settings{
			Title:                  "Loda",
			URL:                    webpackUrl,
			Resizable:              true,
			Width:                  1024,
			Height:                 768,
			ExternalInvokeCallback: handleRPC,
			Debug:                  true,
		})
	default:
		fmt.Printf("Unsupported mode: %v \n", mode)
		os.Exit(0)
		return nil
	}
}

func main() {
	var mode string
	var host string
	var port int
	flag.StringVar(&mode, "mode", "release", "Application mode")
	flag.StringVar(&host, "host", "localhost", "[Debug] Webpack Server Host")
	flag.IntVar(&port, "port", 8080, "[Debug] Webpack Server Port")
	flag.Parse()

	w := createWindow(mode, host, port)
	defer w.Exit()

	ipcMain := ipc.SharedMain()
	ipcMain.SetView(w)

	// Setup Callback

	ipcMain.On(
		"openlink",
		func(event string, value interface{}) interface{} {
			url := value.(string)
			helpers.OpenBrowser(url)
			ipcMain.Send("testing", map[string]string{"testing": "123"})
			return nil
		})

	w.Run()
}
