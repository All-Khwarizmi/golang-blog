package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/All-Khwarizmi/golang-blog/internal/routes"
	"github.com/All-Khwarizmi/golang-blog/internal/web/pages"
	"github.com/gin-gonic/gin"
)

func main() {

	sm := gin.New()
	sm.LoadHTMLGlob("../../internal/web/templates/*")
	sm.GET("/index", pages.Login)
	sm.POST("/login", routes.LoginHandler)

	sm.GET("/home-page", pages.Home)

	srv := &http.Server{
		Addr:         ":9090",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      sm,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
		log.Println("Starting server on port 9090")

	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	sigChan := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	sig := <-sigChan
	log.Println("Recieved terminate, gracefull shutdouwn", sig)

	// The context is used to inform the server it has 30 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting...")
}
