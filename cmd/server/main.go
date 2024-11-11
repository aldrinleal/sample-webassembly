package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stealthrocket/net/wasip1"
	"net/http"
	"os"
)

func getPort() string {
	listener := ":8000"

	if newListenerPort, exists := os.LookupEnv("PORT"); exists {
		listener = ":" + newListenerPort
	}

	return listener
}

func main() {
	e := gin.Default()
	e.Use(gin.Recovery())

	genericHandler := func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	}

	e.GET("/", genericHandler)
	e.GET("/health", genericHandler)
	e.GET("/healthcheck", genericHandler)

	e.Any("/any/*any", genericHandler)

	listenerAddr := getPort()

	listener, err := wasip1.Listen("tcp", listenerAddr)

	if nil != err {
		panic(err)
	}

	server := &http.Server{
		Handler: e,
	}

	if err := server.Serve(listener); nil != err {
		panic(err)
	}
}
