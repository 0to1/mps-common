package logcfg

import (
	"encoding/json"
	"log"
	"runtime"
	"time"

	"github.com/0to1/mps-common/models"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/errors"
	"github.com/sirupsen/logrus"
)

const (
	// FieldFunc ..
	FieldFunc = "func"
	// FieldFile ..
	FieldFile = "file"
	// FieldLine ..
	FieldLine = "line"
)

// DBHook ..
type DBHook struct {
	Service   string
	IP        string
	IsConsole bool
	KDB       *gorm.DB
}

type logContent struct {
	Time   string
	TaskID uint32
	Level  string
	Func   string
	File   string
	Line   int
	Msg    string
}

// Fire ..
func (hook *DBHook) Fire(entry *logrus.Entry) error {
	if hook.KDB == nil {
		return errors.InternalServerError("go.micro.srv.log", "db is invalid")
	}

	l, err := parseEntry(hook, entry)
	if err != nil {
		return err
	}

	err = hook.KDB.Save(&l).Error

	if err != nil {
		log.Println("AddLog error: ", err.Error())
		return err
	}

	return nil
}

// Levels ..
func (hook *DBHook) Levels() []logrus.Level {
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

func parseEntry(hook *DBHook, entry *logrus.Entry) (models.Log, error) {
	var logModel models.Log

	body, err := entry.String()
	if err != nil {
		return models.Log{}, err
	}

	var content logContent

	if err := json.Unmarshal([]byte(body), &content); err != nil {
		return models.Log{}, err
	}

	t, err := time.ParseInLocation("2006-01-02 15:04:05.000", content.Time, time.Local)
	if err != nil {
		return models.Log{}, err
	}

	pc, file, line, ok := runtime.Caller(9)
	if ok {
		funcName := runtime.FuncForPC(pc).Name()

		if len(content.File) < 4 {
			content.File = file
			content.Func = funcName
			content.Line = line
		}
	}

	logModel.Time = t
	logModel.Level = content.Level
	logModel.FuncName = content.Func
	logModel.FileName = content.File
	logModel.Line = content.Line
	logModel.Msg = content.Msg
	logModel.TaskID = content.TaskID
	logModel.Service = hook.Service
	logModel.IP = hook.IP

	if hook.IsConsole {
		log.Println(logModel.Msg, "[", logModel.FileName, " ", logModel.FuncName, " ", logModel.Line, "]")
	}

	return logModel, nil
}
