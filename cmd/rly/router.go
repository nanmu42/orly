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
	"fmt"
	"image/jpeg"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// GenerateCover generate cover and returns it in gif format
func GenerateCover(c *gin.Context) {
	var err error

	cq, err := ParseCoverQuery(c.Request.URL.Query())
	if err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, "%v", err)
		return
	}

	output, err := w.Handle(&cq)
	if err != nil {
		c.Error(err)
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
	// log requests
	router.Use(RequestLogger(logger))

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
		fmt.Println("API starting...")
		logger.Info("API starting...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("API HTTP service: %v", err)
			logger.Fatal("API HTTP service fatal error",
				zap.Error(err),
			)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of `shutdownTimeout` seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, exitSignals...)
	<-quit
	fmt.Println("API is exiting safely...")
	logger.Info("API is exiting safely...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("API exiting timed out:", err)
		logger.Fatal("API exiting timed out",
			zap.Error(err),
		)
	}
	logger.Info("API exited successfully. :)")
	fmt.Println("API exited successfully. :)")

	return
}

// RequestLogger logs every request via zap
func RequestLogger(l *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		receivedAt := time.Now()
		// before
		c.Next()
		// after
		l.Info(c.Request.Method,
			zap.Int("status", c.Writer.Status()),
			zap.String("URL", c.Request.RequestURI),
			zap.String("IP", c.ClientIP()),
			zap.String("UA", c.Request.UserAgent()),
			zap.String("ref", c.Request.Referer()),
			zap.Duration("lapse", time.Now().Sub(receivedAt)),
			zap.Strings("err", c.Errors.Errors()),
		)
	}
}
