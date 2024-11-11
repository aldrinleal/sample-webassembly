package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
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

	run(listenerAddr, e)
}
