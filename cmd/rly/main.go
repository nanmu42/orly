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

	"go.uber.org/zap"

	"github.com/pkg/errors"
)

var (
	configFile = flag.String("config", "config.toml", "config.toml file location for rly")
	w          WorkerPool
	logger     *zap.Logger
)

func main() {
	var err error
	logger, _ = zap.NewProduction()
	defer logger.Sync()

	defer func() {
		if err != nil {
			logger.Fatal("fatal error",
				zap.Error(err),
			)
		}
	}()

	flag.Parse()

	err = C.LoadFrom(*configFile)
	if err != nil {
		err = errors.Wrap(err, "C.LoadFrom")
		return
	}

	err = initializeFactory()
	if err != nil {
		err = errors.Wrap(err, "initializeFactory")
		return
	}

	w = InitWorkerPool(C.WorkerNum, C.QueueLen, makeCover)

	router := setupRouter()
	startAPI(router, C.Port)
}
