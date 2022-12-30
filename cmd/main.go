package main

import (
	Logger "github.com/cjreeder/logging-library"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

type WebServer struct {
	port string
	Log  *zap.SugaredLogger
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
		Log:  L.Log,
	}

	W.buildHTTPServer()
}
