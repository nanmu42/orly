package main

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
