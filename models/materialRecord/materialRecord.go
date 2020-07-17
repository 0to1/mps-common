package materialRecord

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	merrors "github.com/micro/go-micro/v2/errors"
)

// MaterialInCellRecord 物料在储位的记录
type MaterialInCellRecord struct {
	MaterialID int64
	ArriveTime time.Time
	LeaveTime  time.Time
}

// IsExist 物料信息是否存在
func IsExist(db *gorm.DB, materialID int64) (bool, error) {
	var count uint
	err := db.Model(&MaterialInCellRecord{}).Where("material_id = ? and leave_time <= arrive_time", materialID).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// Arrive 入库
func Arrive(db *gorm.DB, materialID int64) error {
	res, err := IsExist(db, materialID)
	if err != nil {
		return err
	}

	if res {
		return merrors.Forbidden("go.micro.srv.material", "materialID is exist")
	}

	err = db.Create(&MaterialInCellRecord{MaterialID: materialID, ArriveTime: time.Now()}).Error
	if err != nil {
		return err
	}
	return nil
}

// Leave 出库
func Leave(db *gorm.DB, materialID int64) error {
	res, err := IsExist(db, materialID)
	if err != nil {
		return err
	}

	if !res {
		return merrors.Forbidden("go.micro.srv.material", "materialID is not exist")
	}

	err = db.Model(&MaterialInCellRecord{}).Where("material_id = ? and leave_time <= arrive_time", materialID).Update("leave_time", time.Now()).Error
	if err != nil {
		return err
	}
	return nil
}

// GetRecord ..
func GetRecord(db *gorm.DB, limit uint32, offset uint32) ([]MaterialInCellRecord, error) {
	var records []MaterialInCellRecord
	err := db.Model(&MaterialInCellRecord{}).Order("arrive_time").Limit(limit).Offset(offset).Find(&records).Error
	if err != nil {
		log.Println("GetRecord error: ", err.Error())
		return nil, err
	}
	return records, nil
}

// GetRecordByDate ..
func GetRecordByDate(db *gorm.DB, startTime string, endTime string) ([]MaterialInCellRecord, error) {
	var records []MaterialInCellRecord
	err := db.Model(&MaterialInCellRecord{}).Where("arrive_time > ? and arrive_time < ?", startTime, endTime).
		Order("arrive_time").Find(&records).Error
	if err != nil {
		log.Println("GetRecord error: ", err.Error())
		return nil, err
	}
	return records, nil
}
