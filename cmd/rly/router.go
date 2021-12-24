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
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/nanmu42/gzip"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// requestCounter is used to count request num
// should only be affected by atomic action
var requestCounter uint64

// GenerateCover generate cover and returns it in gif format
func GenerateCover(c *gin.Context) {
	var err error

	cq, err := ParseCoverQuery(c.Request.URL.Query())
	if err != nil {
		_ = c.Error(err)
		c.String(http.StatusBadRequest, "%v", err)
		return
	}

	output, err := w.Handle(&cq)
	if err != nil {
		_ = c.Error(err)
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}

	c.Set("cq", &cq)

	// set cache control
	c.DataFromReader(http.StatusOK, int64(output.Len()), "image/jpeg", output, map[string]string{})
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
	img, err := factory.Draw(short(t.Query.Title, 150), short(t.Query.TopText, 180), short(t.Query.Author, 150), short(t.Query.GuideText, 150), t.Query.GuideTextPlacement, t.Query.PrimaryColor, int(t.Query.ImageID))
	if err != nil {
		return
	}

	err = jpeg.Encode(&output, img, &jpeg.Options{Quality: 85})
}

func setupRouter() (mux *http.ServeMux) {
	if C.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.HandleMethodNotAllowed = true
	router.Use(gin.Recovery())
	// log requests
	router.Use(RequestLogger(logger))

	// setup routes
	router.GET("/api/generate", GenerateCover)
	router.GET("/api/hello", Hello)

	mux = http.NewServeMux()
	mux.Handle("/api/", router)
	mux.Handle("/", gzip.DefaultHandler().WrapHandler(http.FileServer(gin.Dir("web", false))))

	return
}

// Hello for health check
func Hello(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func startAPI(handler http.Handler, port string) {
	// timeout for safe exit
	const shutdownTimeout = 10 * time.Second

	var exitSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}

	server := &http.Server{
		Addr:    port,
		Handler: handler,
		// ReadTimeout time limit to read the request
		ReadTimeout: 4 * time.Second,
		// WriteTimeout time limit for request reading and response
		WriteTimeout: 8 * time.Second,
		// IdleTimeout keep-alive waiting time
		IdleTimeout: 60 * time.Second,
		// MaxHeaderBytes max header is 8KB
		MaxHeaderBytes: 1 << 13,
	}

	go func() {
		// service connections
		logger.Info("API starting...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
	logger.Info("API is exiting safely...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("API exiting timed out",
			zap.Error(err),
		)
	}
	logger.Info("API exited successfully. :)")
}

// RequestLogger logs every request via zap
func RequestLogger(l *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// request timer
		receivedAt := time.Now()
		// counter++
		sequence := atomic.AddUint64(&requestCounter, 1)

		// before
		c.Next()
		// after
		var cq *CoverQuery
		if payload, exist := c.Get("cq"); exist {
			var ok bool
			if cq, ok = payload.(*CoverQuery); ok && cq != nil {
				l.Info("APIAuditLog",
					zap.Uint64("seq", sequence),
					zap.String("method", c.Request.Method),
					zap.Int("status", c.Writer.Status()),
					zap.String("title", cq.Title),
					zap.String("author", cq.Author),
					zap.String("top", cq.TopText),
					zap.String("guide", cq.GuideText),
					zap.String("guideP", cq.GuideTextPlacement),
					zap.Int64("imageID", cq.ImageID),
					zap.String("IP", c.ClientIP()),
					zap.String("UA", c.Request.UserAgent()),
					zap.String("ref", c.Request.Referer()),
					zap.Duration("lapse", time.Since(receivedAt)),
					zap.Strings("err", c.Errors.Errors()),
				)
				return
			}
		}
	}
}
