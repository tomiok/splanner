package main

const (
	maxQueue   = 10
	maxWorkers = 5
)

var jobQueue chan Unit

type Unit struct {
	job func()
	jobWithErr func() error
}

func initQueue() {
	jobQueue = make(chan Unit, maxQueue)
}

type dispatcher struct {
	pool    chan chan Unit
	workers int
}

type worker struct {
	pool    chan chan Unit
	jobChan chan Unit
	quit    chan bool
}

func newWorker(pool chan chan Unit) *worker {
	return &worker{
		jobChan: make(chan Unit),
		quit:    make(chan bool),
		pool:    pool,
	}
}

func (w *worker) start() {
	go func() {
		for {
			//register the actual worker in the queue
			w.pool <- w.jobChan
			select {
			case job := <-w.jobChan:
				job.job()
			case <-w.quit:
				return
			}
		}
	}()
}

// new dispatcher
func newDispatcher() *dispatcher {
	return &dispatcher{
		pool:    make(chan chan Unit, maxWorkers),
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
		go func(j Unit) {
			jobChannel := <-d.pool
			jobChannel <- j
		}(job)
	}
}
