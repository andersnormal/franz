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

	return nil
}
