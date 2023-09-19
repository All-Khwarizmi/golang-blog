package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ServerConfigWrapper(sm *gin.Engine) *http.Server {
	return &http.Server{
		Addr:         ":9090",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      sm,
	}
}
