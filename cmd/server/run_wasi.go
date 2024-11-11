//go:build wasip1
// +build wasip1

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stealthrocket/net/wasip1"
	"net/http"
)

func run(listenerAddr string, e *gin.Engine) {
	listener, err := wasip1.Listen("tcp", listenerAddr)

	if nil != err {
		logrus.Fatalf("Oops: %s", err)
	}

	logrus.Infof("Going to listen on %s", listenerAddr)

	server := &http.Server{
		Handler: e,
	}

	if err := server.Serve(listener); nil != err {
		logrus.Fatalf("Oops: %s", err)
	}
}
