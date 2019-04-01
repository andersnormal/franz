package pool

import (
	"context"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
)

type worker struct {
	id      int
	work    workQueue
	workers workersQueue

	port int
	r    *runner.Runner
	c    *chromedp.CDP

	exit chan bool
}

func newWorker(id int, workers workersQueue, exit chan bool) (*worker, error) {
	w := new(worker)

	w.id = id
	w.work = make(workQueue)
	w.workers = workers
	w.exit = exit

	opts := []runner.CommandLineOption{
		runner.ExecPath(runner.LookChromeNames("headless_shell")),
		runner.RemoteDebuggingPort(w.port),
		runner.NoDefaultBrowserCheck,
		runner.NoFirstRun,
		runner.Headless,
	}

	r, err := runner.New(opts...)
	if err != nil {
		return nil, err
	}
	w.r = r

	return w, nil
}

func (w *worker) start(ctx context.Context) error {
	err := w.r.Start(ctx)
	if err != nil {
		return err
	}

	w.c, err := chromedp.New(
		ctx,
		chromedp.WithRunner(w.r),
	)
	if err != nil {
		return nil
	}

	for {
		w.workers <- w.work

		select {
		case <-ctx.Done():
			return nil
		case work := <-w.work:
			w.c.Run(ctx, work)
		}
	}
}
