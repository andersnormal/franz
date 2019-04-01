package pool

// WithSize ...
func WithSize(size int) func(o *Opts) {
	return func(o *Opts) {
		o.Size = size
	}
}

// WithBuffer ...
func WithBuffer(buffer int) func(o *Opts) {
	return func(o *Opts) {
		o.Buffer = buffer
	}
}

func configure(p *pool, opts ...Opt) error {
	for _, o := range opts {
		o(p.opts)
	}

	p.size = p.opts.Size
	p.buffer = p.opts.Buffer
	p.end = p.opts.End
	p.start = p.opts.Start

	return nil
}
