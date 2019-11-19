package logcfg

import "github.com/sirupsen/logrus"

type MicroHook struct {
	Service string
	IP      string
}

func (hook *MicroHook) Fire(entry *logrus.Entry) error {
	entry.Data["service"] = hook.Service
	entry.Data["ip"] = hook.IP
	return nil
}

func (hook *MicroHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.TraceLevel,
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}
