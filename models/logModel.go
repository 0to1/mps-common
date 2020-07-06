package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// Log ..
type Log struct {
	Time     time.Time
	TaskID   uint32
	Level    string
	FuncName string
	FileName string
	Line     int
	Service  string
	IP       string
	Msg      string //日志详情
}

// AddLog ..
func AddLog(db *gorm.DB, wTime time.Time, level string, funcName string, line int,
	fileName string, srvName string, ip string, msg string, taskID uint32) error {

	var logDB Log
	logDB.Time = wTime
	logDB.Level = level
	logDB.FuncName = funcName
	logDB.FileName = fileName
	logDB.Service = srvName
	logDB.IP = ip
	logDB.Msg = msg
	logDB.Line = line
	logDB.TaskID = taskID

	err := db.Save(&logDB).Error

	if err != nil {
		log.Println("AddLog error: ", err.Error())
		return err
	}

	return nil
}
