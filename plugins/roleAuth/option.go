package roleauth

import "container/list"

type Options struct {
	serviceName string
	skipper     *list.List
}

type Option func(*Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		serviceName: "go.micro.srv.user",
		skipper:     list.New(),
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func ServiceName(name string) Option {
	return func(o *Options) {
		o.serviceName = name
	}
}

func Skipper(skipper []string) Option {
	return func(o *Options) {
		for _, url := range skipper {
			// 跳过已经存在的url
			for v := o.skipper.Front(); v != nil; v = v.Next() {
				if v.Value == url {
					continue
				}
			}

			o.skipper.PushBack(url)
		}
	}
}
