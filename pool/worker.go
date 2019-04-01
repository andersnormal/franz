package pool

type worker struct {
	id      int
	work    workQueue
	workers workersQueue

	failed chan int
}

func newWorker(id int, workers workersQueue, failed chan int) *worker {
	w := new(worker)

	w.id = id
	w.work = make(workQueue)
	w.workers = workers
	w.failed = failed

	return w
}

func (w *worker) start() {}
