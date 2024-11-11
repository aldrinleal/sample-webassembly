//go:build !wasm
// +build !wasm

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func run(listenerAddr string, e *gin.Engine) {
	logrus.Infof("Going to listen on %s", listenerAddr)

	logrus.Fatalf("Oops: %s", http.ListenAndServe(listenerAddr, e))
}
