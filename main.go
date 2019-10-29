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

	webview "github.com/darkcl/loda/lib/webview"
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

func bundleAssets(dir string) (string, string) {
	js := mewn.String("./ui/dist/bundle.min.js")
	indexHTML := mewn.String("./ui/dist/index.html")

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

	return abs, tmpJs
}

func createWindow(mode string, host string, port int, dir string) webview.WebView {
	switch mode {
	case "release":
		fmt.Println("Release Mode")
		indexPath, _ := bundleAssets(dir)

		return webview.New(webview.Settings{
			Title:                  "Loda",
			URL:                    "file://" + indexPath,
			Resizable:              true,
			Width:                  1024,
			Height:                 768,
			ExternalInvokeCallback: handleRPC,
			Debug:                  true,
		})
	case "debug":
		fmt.Println("Debug Mode")
		webpackURL := fmt.Sprintf("http://%s:%v", host, port)
		return webview.New(webview.Settings{
			Title:                  "Loda",
			URL:                    webpackURL,
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

	dir, err := ioutil.TempDir("", "loda")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer os.RemoveAll(dir)

	w := createWindow(mode, host, port, dir)
	defer w.Exit()

	ipcMain := ipc.SharedMain()
	ipcMain.SetView(w)

	// Setup Callback

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

	w.Run()
}
