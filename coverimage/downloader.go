// +build downloader

/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

// AssetManager downloads file via URL provided in `asset.list`
// and rename it in one-based sequence number in list.
//
// This tool only relies on file name for existing check.
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/pkg/errors"
)

const (
	listName    = "asset.list"
	concurrency = 4
)

var (
	processed int64
	wg        sync.WaitGroup
)

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Printf("%v\n", err)
			os.Exit(1)
		}
	}()

	list, err := readList(listName)
	if err != nil {
		err = errors.Wrap(err, "readList")
		return
	}

	downloaded, err := readDir()
	if err != nil {
		err = errors.Wrap(err, "readDir")
		return
	}

	log.Printf("downloading %v files with concurrency %v...\n", len(list), concurrency)

	wg.Add(len(list))
	var backet = make(chan bool, concurrency)
	for index, task := range list {
		backet <- true
		go download(task, fmt.Sprintf("%v.tif", index), downloaded, backet)
	}

	wg.Wait()
	log.Printf("download complete: %v/%v\n", processed, len(list))
}

func readDir() (downloaded *sync.Map, err error) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		err = errors.Wrap(err, "ioutil.ReadDir")
		return
	}

	downloaded = new(sync.Map)
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		downloaded.Store(f.Name(), true)
	}
	return
}

func readList(fileName string) (URLs []string, err error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		err = errors.Wrap(err, "ioutil.ReadFile")
		return
	}
	// list is divided by return
	pieces := strings.Split(string(content), "\n")
	URLs = make([]string, 0, len(pieces))
	// comments are started by #
	for _, item := range pieces {
		item = strings.Trim(item, " \r")
		if len(item) == 0 || strings.HasPrefix(item, "#") {
			continue
		}
		URLs = append(URLs, item)
	}
	return
}

func download(URLs string, name string, downloaded *sync.Map, done <-chan bool) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("[%s] %v", name, err)
		} else {
			atomic.AddInt64(&processed, 1)
		}

		wg.Done()
		<-done
	}()
	// file exist?
	if _, ok := downloaded.Load(name); ok {
		log.Printf("[%s] already exists, skipped\n", name)
		return
	}

	resp, err := http.Get(URLs)
	if err != nil {
		err = errors.Wrap(err, "http.Get")
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(name)
	if err != nil {
		err = errors.Wrap(err, "os.Create")
		return
	}
	defer func() {
		if err != nil {
			os.Remove(name)
		}
	}()
	defer f.Close()
	buffered := bufio.NewWriterSize(f, 1024*1024)

	n, err := io.Copy(buffered, resp.Body)
	if err != nil {
		err = errors.Wrap(err, "io.Copy")
		return
	}
	err = buffered.Flush()
	if err != nil {
		err = errors.Wrap(err, "buffered.Flush()")
		return
	}

	log.Printf("[%s] successfully downloaded %v KB\n", name, n/1024)
	return
}
