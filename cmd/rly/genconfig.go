// +build ignore

/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	C.QueueLen = 10
	C.WorkerNum = 4

	content, err := C.Info()
	if err != nil {
		err = errors.Wrap(err, "C.Info")
		return
	}
	f, err := os.Create("config_example.toml")
	if err != nil {
		err = errors.Wrap(err, "os.Create")
		return
	}
	defer f.Close()

	content = `# Copy this file to create default config
#
# cp config_example.toml config.toml

` + content

	_, err = f.WriteString(content)
	if err != nil {
		err = errors.Wrap(err, "f.WriteString")
		return
	}

	return
}
