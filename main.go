package main

import (
	"flag"

	"github.com/darkcl/loda/application"
)

func main() {
	var mode string
	var host string
	var port string
	flag.StringVar(&mode, "mode", "release", "Application mode")
	flag.StringVar(&host, "host", "localhost", "[Debug] Webpack Server Host")
	flag.StringVar(&port, "port", "8080", "[Debug] Webpack Server Port")
	flag.Parse()

	app := &application.DesktopApplication{}

	defer app.WillTerminate()

	app.WillLaunch(mode, map[string]string{
		"host": host,
		"port": port,
	})

	app.DidFinishLaunching()
}
