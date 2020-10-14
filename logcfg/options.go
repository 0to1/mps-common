package logcfg

import "github.com/sirupsen/logrus"

// Options 选项
type Options struct {
	Hooks           []logrus.Hook
	ReportCaller bool
	Level        logrus.Level
}

type Option func(*Options)

func newOptions(opts ...Option) Options {
	options := Options{
		Hooks: make([]logrus.Hook, 0),
		ReportCaller: false,
		Level: logrus.InfoLevel,
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

//SetReportCaller 是否允许打印函数调用信息
func SetReportCaller(isstart bool) Option {
	return func(o *Options) {
		o.ReportCaller = isstart
	}
}
//Setlevel 设置默认日志级别
func Setlevel(level logrus.Level) Option {
	return func(o *Options) {
		o.Level = level
	}
}
