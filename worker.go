package splanner

const (
	maxQueue   = 10
	maxWorkers = 5
)

var JobQueue chan Unit

type Unit struct {
	Job        func()
	jobWithErr func() error
}

func InitQueue() {
	JobQueue = make(chan Unit, maxQueue)
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

func (w *worker) Start() {
	go func() {
		for {
			//register the actual worker in the queue
			w.pool <- w.jobChan
			select {
			case job := <-w.jobChan:
				job.Job()
			case <-w.quit:
				return
			}
		}
	}()
}

// NewDispatcher create a pointer to a dispatcher struct
func NewDispatcher() *dispatcher {
	return &dispatcher{
		pool:    make(chan chan Unit, maxWorkers),
		workers: maxWorkers,
	}
}

func (d *dispatcher) Run() {
	for i := 0; i < d.workers; i++ {
		w := newWorker(d.pool)
		w.Start()
	}
	go d.dispatch()
}

func (d *dispatcher) dispatch() {
	for job := range JobQueue {
		go func(j Unit) {
			jobChannel := <-d.pool
			jobChannel <- j
		}(job)
	}
}
