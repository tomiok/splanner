package splanner

var JobQueue chan Unit
var QuitChan chan bool

type Unit struct {
	Job func()
}

func InitQueue(maxQueue int) {
	JobQueue = make(chan Unit, maxQueue)
	QuitChan = make(chan bool, 1)
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

func (w *worker) Start() {
	go func() {
		for {
			// register the actual worker in the queue.
			w.pool <- w.jobCh
			select {
			case job := <-w.jobCh:
				// do the actual job here
				job.Job()
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

func (d *dispatcher) Run(async bool) {
	for i := 0; i < d.workers; i++ {
		w := newWorker(d.pool)
		w.Start()
	}
	if async {
		go d.dispatchAsync()
	} else {
		go d.dispatch()
	}
}

func (d *dispatcher) dispatchAsync() {
	for job := range JobQueue {
		go func(j Unit) {
			jobChannel := <-d.pool
			jobChannel <- j
		}(job)
	}
}

// not async
func (d *dispatcher) dispatch() {
	go func() {
		for {
			select {
			case job, ok := <-JobQueue:
				if ok {
					jobChannel := <-d.pool
					jobChannel <- job
				}
			}
		}
	}()
}
