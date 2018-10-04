package main

import (
	"bytes"
	"errors"
)

// ErrBusyService the queue is full
var ErrBusyService = errors.New("service is busy, please try again later")

// Task carries the request and encoded result
type Task struct {
	Query      *CoverQuery
	ResultSlot chan *TaskResult
}

// TaskResult is a holder for result
type TaskResult struct {
	EncodedImg *bytes.Buffer
	Err        error
}

// WorkerPool is a place for workers to get their jobs.
type WorkerPool struct {
	workerNum  int
	queueLen   int
	inputChan  chan *Task
	handleFunc func(*Task)
}

// InitWorkerPool register handleFunc and runs worker due to workerNum
func InitWorkerPool(workerNum, queueLen int, handleFunc func(inputChan *Task)) (w WorkerPool) {
	w.workerNum, w.queueLen, w.handleFunc = workerNum, queueLen, handleFunc
	w.inputChan = make(chan *Task, w.workerNum*w.queueLen)
	for i := 0; i < workerNum; i++ {
		go w.newWorker()
	}
	return
}

func (w *WorkerPool) newWorker() {
	for t := range w.inputChan {
		w.handleFunc(t)
	}
}

// Handle accomplish query via workers
func (w *WorkerPool) Handle(query *CoverQuery) (*bytes.Buffer, error) {
	t := Task{
		Query:      query,
		ResultSlot: make(chan *TaskResult, 1),
	}
	select {
	case w.inputChan <- &t:
		// relax
	default:
		// the queue is full
		return nil, ErrBusyService
	}
	result := <-t.ResultSlot
	return result.EncodedImg, result.Err
}
