package main

import (
	Logger "github.com/cjreeder/logging-library"
	"github.com/spf13/pflag"
)

type WebServer struct {
	port string
	Log  Logger.Logger
}

func main() {
	var port string
	pflag.StringVarP(&port, "port", "p", "8013", "port for microservice to av-api communication")
	pflag.Parse()
	port = ":" + port

	// Initiate the logger
	L := Logger.NewLogger("info")
	L.Log.Infof("Starting the webserver......")

	// Initiate the WebServer
	W := &WebServer{
		port: port,
		Log:  Logger.Logger{},
	}
	W.BuildHTTPServer()
}
