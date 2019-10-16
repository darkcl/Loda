package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/darkcl/skeleton-go-desktop/helpers"
	"github.com/darkcl/skeleton-go-desktop/ipc"

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

func main() {
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

	w := webview.New(webview.Settings{
		Title:                  "Skeleton",
		URL:                    "file://" + abs,
		Resizable:              true,
		Width:                  1024,
		Height:                 768,
		ExternalInvokeCallback: handleRPC,
		Debug:                  true,
	})
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
