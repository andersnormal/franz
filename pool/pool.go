package pool

import (
	"net"
	"strings"
)

func New(opts ...Opt) (Pool, error) {
	options := &Opts{}

	p := new(pool)
	p.opts = options

	configure(p, opts...)

	p.failed = make(chan int, 1)
	p.work = make(workQueue, p.buffer)
	p.workersQueue = make(workersQueue, p.size)

	for i := 0; i < p.size; i++ {
		worker := newWorker(i, p.workersQueue, p.failed)
		p.workers = append(p.workers, worker)

		port := p.start + i
		if port >= p.end {
			return p, nil
		}

		if err := testListener(p.start + i); err != nil {
			return nil, err
		}

		go worker.start()
	}

	return p, nil
}

func (p *pool) Len() int {
	return 0
}

func (p *pool) Run() error {
	return nil
}

func (p *pool) Shutdown() error {

	return nil
}

func (p *pool) watch() {
	// todo: add interrupt

	for {
		select {
		case <-p.failed:
			// start a new instance
			i := len(p.workers)
			worker := newWorker(i, p.workersQueue, p.failed)
			p.workers = append(p.workers, worker)

			go worker.start()
		}
	}
}

func testListener(port int) error {
	l, err := net.Listen("tcp", strings.Join([]string{"localhost", string(port)}, ":"))
	if err != nil {
		return err
	}
	defer l.Close()

	return nil
}

// // Queue is the interface to a queue itself
// type Queue interface {
// 	// Add is adding something to the queue to be processed.
// 	Add(interface{})
// 	// Wait is waiting for the queued work to be finished,
// 	// or getting free space in the queue.
// 	Wait() chan Queueable
// 	// Len is retuning the current length of work in the queue.
// 	Len() int
// }

// // WorkerQueue additionally adds the interface to the workers for
// // doing the work in the queue that is worked by workers
// type WorkerQueue interface {
// 	// Shutdown is shutting done the underlying workers.
// 	Shutdown()

// 	Queue
// }

// // WorkFunc is a callback function provided to the workers
// // which actually process the work that is coming in.
// type WorkFunc func(queueable interface{}) (interface{}, error)

// // Queueable ...
// type Queueable struct {
// 	value interface{}
// 	err   error
// }

// // Value ...
// func (q *Queueable) Value() interface{} {
// 	return q.value
// }

// // Err ...
// func (q *Queueable) Err() error {
// 	return q.err
// }

// // Opt ...
// type Opt func(*Opts)

// // Opts ...
// type Opts struct {
// 	Size     int
// 	Workers  int
// 	WorkFunc WorkFunc
// }

// NewWorkerQueue creates a new queue for workers.
//
//	q := NewWorkerQueue(func(o *Opts) {
//		o.Size = 1
//		o.Workers = 1
//		o.WorkFunc = fn
//	})
//
// func NewWorkerQueue(opts ...Opt) WorkerQueue {
// 	options := &Opts{}

// 	q := new(queue)
// 	q.opts = options

// 	q.workersQueue = make(workersQueue, q.num)
// 	q.work = make(workQueue, q.size)
// 	q.finish = make(finishQueue, q.size)

// 	configure(q, opts...)

// 	for i := 0; i < q.num; i++ {
// 		worker := newWorker(i, q.workersQueue, q.finish, q.fn)
// 		q.workers = append(q.workers, worker)

// 		go worker.start()
// 	}

// 	// currate pipeping work to the workers
// 	go q.watch()

// 	return q
// }

// type workQueue chan interface{}
// type workersQueue chan chan interface{}
// type finishQueue chan Queueable

// type worker struct {
// 	id      int
// 	fn      WorkFunc
// 	work    workQueue
// 	workers workersQueue
// 	finish  finishQueue

// 	quit chan bool

// 	wg *sync.WaitGroup
// }

// type queue struct {
// 	work         workQueue
// 	workersQueue workersQueue

// 	finish finishQueue

// 	counter int
// 	workers []*worker

// 	size int
// 	num  int

// 	fn WorkFunc

// 	opts *Opts

// 	once sync.Once

// 	sync.RWMutex
// }

// func newWorker(id int, workers workersQueue, finish finishQueue, fn WorkFunc) *worker {
// 	w := &worker{
// 		id:      id,
// 		work:    make(workQueue),
// 		workers: workers,
// 		quit:    make(chan bool),
// 		finish:  finish,
// 		fn:      fn,
// 	}

// 	return w
// }

// func (q *queue) watch() func() {
// 	for {
// 		select {
// 		case work := <-q.work:
// 			go func() {
// 				worker := <-q.workersQueue
// 				worker <- work
// 			}()
// 		default:
// 			break
// 		}
// 	}
// }

// // Len ...
// func (q *queue) Len() int {
// 	return len(q.work)
// }

// // Add ...
// func (q *queue) Add(v interface{}) {
// 	q.work <- v
// }

// // Shutdown ...
// func (q *queue) Shutdown() {
// 	q.once.Do(func() {
// 		for _, worker := range q.workers {
// 			worker.stop()
// 		}
// 	})
// }

// // Get ...
// func (q *queue) Get() Queueable {
// 	v := <-q.finish

// 	return v
// }

// // Wait ...
// func (q *queue) Wait() chan Queueable {
// 	return q.finish
// }

// func (w *worker) start() {
// 	for {
// 		w.workers <- w.work

// 		select {
// 		case <-w.quit:
// 			return
// 		case work := <-w.work:
// 			v, err := w.fn(work)
// 			w.finish <- Queueable{value: v, err: err}
// 		}
// 	}
// }

// func (w *worker) stop() {
// 	go func() {
// 		w.quit <- true
// 	}()
// }

// func configure(q *queue, opts ...Opt) error {
// 	for _, o := range opts {
// 		o(q.opts)
// 	}

// 	q.size = q.opts.Size
// 	q.num = q.opts.Workers
// 	q.fn = q.opts.WorkFunc

// 	return nil
// }
