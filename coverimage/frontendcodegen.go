// +build codegen

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
	"strings"
)

const max = 40

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	b := strings.Builder{}
	for i := 0; i <= max; i++ {
		b.WriteString(fmt.Sprintf(`<div class="animal"><img src="../assets/thumbnails/%v.tif.gif" alt="%v"/>%v</div>
`, i, i, i))
	}
	fmt.Println(b.String())
}
