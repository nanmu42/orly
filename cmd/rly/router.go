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

	output, err := w.Handle(&cq)
	if err != nil {
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}

	// set cache control
	c.DataFromReader(http.StatusOK, int64(output.Len()), "image/jpeg", output, map[string]string{
		"Cache-Control": "public, max-age=86400",
	})
}

func makeCover(t *Task) {
	var (
		output bytes.Buffer
		err    error
	)
	defer func() {
		t.ResultSlot <- &TaskResult{
			EncodedImg: &output,
			Err:        err,
		}
	}()

	if t.Query.ImageID > int64(C.MaxImageID) || t.Query.ImageID < 0 {
		t.Query.ImageID = 0
	}
	img, err := factory.Draw(short(t.Query.Title, 42), short(t.Query.TopText, 60), short(t.Query.Author, 36), short(t.Query.GuideText, 40), t.Query.GuideTextPlacement, t.Query.PrimaryColor, int(t.Query.ImageID))
	if err != nil {
		return
	}

	err = jpeg.Encode(&output, img, &jpeg.Options{85})

	return
}

func setupRouter() (router *gin.Engine) {
	if C.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router = gin.New()
	router.HandleMethodNotAllowed = true
	router.Use(gin.Recovery())
	// set debug mode if in need
	if C.Debug {
		router.Use(gin.Logger())
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
		log.Println("API starting...")
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
