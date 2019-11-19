package cfghelper

type Options struct {
	source string
}

type Option func(*Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		source: "none",
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func Source(source string) Option {
	return func(o *Options) {
		o.source = source
	}
}
