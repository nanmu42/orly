/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

package main

import (
	"bytes"
	"context"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// GenerateCover generate cover and returns it in gif format
func GenerateCover(c *gin.Context) {
	var err error

	cq, err := ParseCoverQuery(c.Request.URL.Query())
	if err != nil {
		c.String(http.StatusBadRequest, "%v", err)
		return
	}

	if cq.ImageID > int64(C.MaxImageID) || cq.ImageID < 0 {
		cq.ImageID = 0
	}
	img, err := factory.Draw(short(cq.Title, 42), short(cq.TopText, 60), short(cq.Author, 36), short(cq.GuideText, 40), cq.GuideTextPlacement, cq.PrimaryColor, int(cq.ImageID))
	if err != nil {
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}

	var output bytes.Buffer
	err = jpeg.Encode(&output, img, &jpeg.Options{85})
	if err != nil {
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}

	// set cache control
	c.Header("Cache-Control", "public, max-age=86400")
	c.Data(http.StatusOK, "image/jpeg", output.Bytes())
}

func setupRouter() (router *gin.Engine) {
	router = gin.New()
	router.HandleMethodNotAllowed = true
	router.Use(gin.Recovery())
	// set debug mode if in need
	if C.Debug {
		gin.SetMode(gin.DebugMode)
		router.Use(gin.Logger())
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// setup routes
	router.GET("/generate", GenerateCover)
	return
}

func startAPI(handler http.Handler, port string) {
	// timeout for safe exit
	const shutdownTimeout = 2 * time.Minute

	var exitSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}

	server := &http.Server{
		Addr:    port,
		Handler: handler,
		// ReadTimeout time limit to read the request
		ReadTimeout: 10 * time.Second,
		// WriteTimeout time limit for request reading and response
		WriteTimeout: 40 * time.Second,
		// IdleTimeout keep-alive waiting time
		IdleTimeout: 60 * time.Second,
		// MaxHeaderBytes max header is 8KB
		MaxHeaderBytes: 1 << 13,
	}

	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("API HTTP Service: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of `shutdownTimeout` seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, exitSignals...)
	<-quit
	log.Println("API is exiting safely...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("API exiting timed out:", err)
	}
	log.Println("API exited successfully. :)")

	return
}
