package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stealthrocket/net/wasip1"
	ginlogrus "github.com/toorop/gin-logrus"
	"net/http"
	"os"
	"strings"
)

func getPort() string {
	listener := ":8000"

	if newListenerPort, exists := os.LookupEnv("PORT"); exists {
		listener = ":" + newListenerPort
	}

	return listener
}

var log = logrus.New()

func main() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)

	e := gin.Default()
	e.Use(ginlogrus.Logger(log), gin.Recovery())

	genericHandler := func(c *gin.Context) {
		log.Infof("Request for %s", c.Request.URL)
		for k, v := range c.Request.Header {
			log.Infof(" %s: %s", k, strings.Join(v, "; "))
		}

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
		log.Fatalf("Oops: %s", err)
	}

	log.Infof("Going to listen on %s", listenerAddr)

	server := &http.Server{
		Handler: e,
	}

	if err := server.Serve(listener); nil != err {
		log.Fatalf("Oops: %s", err)
	}
}
