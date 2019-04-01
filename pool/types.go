package pool

import ()

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	Buffer int
	Size   int

	Start int
	End   int
}

// Pool ...
type Pool interface {
	// Run is running work on a worker
	Run() error
	// Shutdown is gracefully shutting down the pool
	Shutdown() error
	// Len is returning the current size of the queue to the workers pool
	Len() int
}

type pool struct {
	buffer int
	size   int

	start int
	end   int

	work         workQueue
	workersQueue workersQueue

	workers []*worker

	failed chan int

	opts *Opts
}

type workQueue chan interface{}
type workersQueue chan chan interface{}
