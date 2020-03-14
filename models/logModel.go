package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// LogModel ..
type LogModel struct {
	Time     time.Time
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
	fileName string, srvName string, ip string, msg string) error {

	var logDB LogModel
	logDB.Time = wTime
	logDB.Level = level
	logDB.FuncName = funcName
	logDB.FileName = fileName
	logDB.Service = srvName
	logDB.IP = ip
	logDB.Msg = msg
	logDB.Line = line

	err := db.Save(&logDB).Error

	if err != nil {
		log.Println("AddLog error: ", err.Error())
		return err
	}

	return nil
}
