package splanner

import "log"

var jobQueue chan Unit

type Unit interface {
	Job() error
}

func InitQueue(maxQueue int) {
	jobQueue = make(chan Unit, maxQueue)
}

type dispatcher struct {
	pool    chan chan Unit
	workers int
}

type worker struct {
	pool  chan chan Unit
	jobCh chan Unit
}

func newWorker(pool chan chan Unit) *worker {
	return &worker{
		jobCh: make(chan Unit),
		pool:  pool,
	}
}

// start, ok worker is working now.
func (w *worker) start() {
	go func() {
		for {
			// register the actual worker in the queue.
			w.pool <- w.jobCh
			select {
			case job := <-w.jobCh:
				// do the actual job here
				err := job.Job()

				if err != nil {
					log.Println(err.Error())
				}
			}
		}
	}()
}

// NewDispatcher create a pointer to a dispatcher struct
func NewDispatcher(maxWorkers int) *dispatcher {
	return &dispatcher{
		pool:    make(chan chan Unit, maxWorkers),
		workers: maxWorkers,
	}
}

// Run is the starting point. This should be called by the client.
func (d *dispatcher) Run(async bool) {
	for i := 0; i < d.workers; i++ {
		w := newWorker(d.pool)
		w.start()
	}
	if async {
		go d.dispatchAsync()
	} else {
		go d.dispatch()
	}
}

// dispatchAsync
func (d *dispatcher) dispatchAsync() {
	for job := range jobQueue {
		go func(j Unit) {
			jobChannel := <-d.pool
			jobChannel <- j
		}(job)
	}
}

// dispatch not async
func (d *dispatcher) dispatch() {
	go func() {
		for {
			select {
			case job, ok := <-jobQueue:
				if ok {
					jobChannel := <-d.pool
					jobChannel <- job
				}
			}
		}
	}()
}

func AddUnit(u Unit) {
	jobQueue <- u
}
