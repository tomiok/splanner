package main

import "fmt"

var jobQueue chan Job

const (
	maxQueue   = 10
	maxWorkers = 5
)

func initQueue() {
	jobQueue = make(chan Job, maxQueue)
}

type dispatcher struct {
	pool    chan chan Job
	workers int
}

type worker struct {
	pool    chan chan Job
	jobChan chan Job
	quit    chan bool
}

func newWorker(pool chan chan Job) *worker {
	return &worker{
		jobChan: make(chan Job),
		quit:    make(chan bool),
		pool:    pool,
	}
}

func (w *worker) start() {
	go func() {
		for {

			//register the actual worker
			w.pool <- w.jobChan
			fmt.Println("job worker registered in queue")
			select {
			case job := <-w.jobChan:
				job.P.Run()
			case <-w.quit:
				return
			}
		}
	}()
}

// new dispatcher
func newDispatcher() *dispatcher {
	return &dispatcher{
		pool:    make(chan chan Job, maxWorkers),
		workers: maxWorkers,
	}
}

func (d *dispatcher) run() {
	for i := 0; i < d.workers; i++ {
		worker := newWorker(d.pool)
		worker.start()
	}
	go d.dispatch()
}

func (d *dispatcher) dispatch() {
	for job := range jobQueue {
		go func(j Job) {
			jobChannel := <-d.pool
			jobChannel <- j
		}(job)
	}
}
