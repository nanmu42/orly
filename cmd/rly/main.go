/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

// rly is an API for O'RLY cover generation
package main

import (
	"flag"
	"fmt"

	"go.uber.org/zap"

	"github.com/pkg/errors"
)

var (
	configFile = flag.String("config", "config.toml", "config.toml file location for rly")
	w          WorkerPool
	logger     *zap.Logger
	// Version build params
	Version string
	// BuildDate build params
	BuildDate string
)

func main() {
	flag.Parse()

	var err error
	defer func() {
		if err != nil {
			if logger != nil {
				logger.Error("rly exits with error", zap.Error(err))
			} else {
				fmt.Println(err)
			}
		}
	}()

	err = C.LoadFrom(*configFile)
	if err != nil {
		err = fmt.Errorf("loading configFile file: toml.DecodeFile: %w", err)
		return
	}

	if C.Debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		err = fmt.Errorf("init zap logger: %w", err)
		return
	}
	defer logger.Sync() // nolint: errcheck
	zap.ReplaceGlobals(logger)

	logger.Info(`O'rly Generator API starting...`,
		zap.String("version", Version),
		zap.String("buildAt", BuildDate))

	err = initializeFactory()
	if err != nil {
		err = errors.Wrap(err, "initializeFactory")
		return
	}

	w = InitWorkerPool(C.WorkerNum, C.QueueLen, makeCover)

	router := setupRouter()
	startAPI(router, C.Port)
}
