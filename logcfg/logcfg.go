package logcfg

import (
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

var (
	// LogFileFormat 日志文件名格式
	LogDir        = "logs"
	LogFileFormat = "YF.MPS-%Y%m%d%H%M.log"
	RecentLogFile = "YF.MPS-recent.log"
	Formatter     = "json"
)

// type Fields map[string]interface{}

// Config 日志配置
func Config(opts ...Option) error {
	options := newOptions(opts...)

	if Formatter == "json" {
		// 配置格式为json格式
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05.000",
			// FullTimestamp: true, //text专用
			PrettyPrint: true, //json专用
		})
	} else {
		// 配置格式为文本格式
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05.000",
			FullTimestamp:   true, //text专用
		})
	}

	// 加载钩子函数
	for _, v := range options.Hooks {
		logrus.AddHook(v)
	}

	// 设置是否允许打印函数调用信息
	if options.ReportCaller == true {
		logrus.SetReportCaller(true)
	} 

	// 设置默认日志级别
	if options.Level != logrus.InfoLevel{
		logrus.SetLevel(options.Level)
	}

	dir := createLogDir()

	// 日志分割配置
	logf, err := rotatelogs.New(
		path.Join(dir, LogFileFormat),
		rotatelogs.WithLinkName(path.Join(dir, RecentLogFile)),
		//rotatelogs.WithMaxAge(24*time.Hour*30 /*day*/),
		rotatelogs.WithMaxAge(-1),
		//rotatelogs.WithRotationTime(1*time.Hour),
		rotatelogs.WithRotationTime(5*time.Minute), //5min一个文件
		rotatelogs.WithRotationCount(1000),         //保留1000个文件
	)

	if err != nil {
		return err
	}

	logrus.SetOutput(logf)

	return nil
}

func createLogDir() string {
	curDir, err := os.Getwd()
	if err != nil {
		return LogDir
	}

	logdir := path.Join(curDir, LogDir)

	os.MkdirAll(logdir, os.ModePerm)

	return logdir
}

// GetLogLevel 获取日志等级
func GetLogLevel() logrus.Level {
	return logrus.GetLevel()
}

//SetLogLevel 设置日志等级
func SetLogLevel(level logrus.Level) {
	logrus.SetLevel(level)
}