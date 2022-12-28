package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cjreeder/microservice_test/handlers"
	"github.com/spf13/pflag"
)

func main() {
	var port string
	pflag.StringVarP(&port, "port", "p", "8013", "port for microservice to av-api communication")
	pflag.Parse()
	port = ":" + port

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "good",
		})
	})
	// :address is the address to the device that you want to manage
	// group your api's by version so you can roll out a newer version
	// without breaking backwards compatibility.
	//
	//
	// action endpoints
	route := router.Group("/api/v1")
	route.GET("/:address/power/:power", handlers.SetPower)
	route.GET("/:address/volume/:volume", handlers.SetVolume)
	route.GET("/:address/volume/mute/:mute", handlers.SetMute)
	route.GET("/:address/display/:blank", handlers.SetBlank)
	route.GET("/:address/input/:input", handlers.SetInput)

	// status endpoints
	route.GET("/:address/power", handlers.GetPower)
	route.GET("/:address/volume", handlers.GetVolume)
	route.GET("/:address/volume/mute", handlers.GetMute)
	route.GET("/:address/input", handlers.GetInput)
	route.GET("/:address/booted", handlers.GetBooted)

	server := &http.Server{
		Addr:           port,
		MaxHeaderBytes: 1021 * 10,
	}

	router.Run(server.Addr)
}
