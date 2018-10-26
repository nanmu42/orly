/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

package common

import (
	"bufio"

	"gopkg.in/natefinch/lumberjack.v2"
)

// BufferedLumberjack implements zapcore.WriteSyncer
type BufferedLumberjack struct {
	buffer *bufio.Writer
}

// NewBufferedLumberjack factory of BufferedLumberjack
func NewBufferedLumberjack(lumber *lumberjack.Logger, size int) *BufferedLumberjack {
	return &BufferedLumberjack{
		buffer: bufio.NewWriterSize(lumber, size),
	}
}

// Write implements zapcore.WriteSyncer
func (b *BufferedLumberjack) Write(p []byte) (n int, err error) {
	return b.buffer.Write(p)
}

// Sync implements zapcore.WriteSyncer
func (b *BufferedLumberjack) Sync() error {
	return b.buffer.Flush()
}
