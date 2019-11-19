package logcfg

import "github.com/sirupsen/logrus"

// Options 选项
type Options struct {
	Hooks []logrus.Hook
}

type Option func(*Options)

func newOptions(opts ...Option) Options {
	options := Options{
		Hooks: make([]logrus.Hook, 0),
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

// AddHook 添加hook钩子函数
func AddHook(hook logrus.Hook) Option {
	return func(o *Options) {
		o.Hooks = append(o.Hooks, hook)
	}
}

// AddHooks 添加多个hook钩子函数
func AddHooks(hooks []logrus.Hook) Option {
	return func(o *Options) {
		o.Hooks = hooks
	}
}
