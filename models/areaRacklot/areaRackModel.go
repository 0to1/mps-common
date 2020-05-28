package areaRacklot

import (
	"errors"

	"github.com/jinzhu/gorm"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

// AreaRacklotModel ..
type AreaRacklot struct {
	gorm.Model
	AreaID    int
	RacklotID int
}

func Query(db *gorm.DB, AreaID int, RacklotID int) (*AreaRacklot, error) {

	if AreaID == 0 && RacklotID == 0 {
		return nil, nil
	}

	var m AreaRacklot
	db = db.Where("area_id = ? and racklot_id = ?", AreaID, RacklotID).First(&m)
	if err := db.Error; err != nil {
		return nil, err
	}

	if db.RowsAffected <= 0 {
		return nil, nil
	}

	return &m, nil
}

func IsExit(db *gorm.DB, AreaID int, RacklotID int) (bool, error) {
	if AreaID == 0 && RacklotID == 0 {
		return false, nil
	}

	db = db.Where("area_id = ? and racklot_id = ?", AreaID, RacklotID).First(&AreaRacklot{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func IsExitByRacklot(db *gorm.DB, racklotID int) (bool, error) {
	if racklotID == 0 {
		return false, nil
	}

	db = db.Where("racklot_id  = ? ", racklotID).First(&AreaRacklot{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}
	return true, nil
}

func IsExitByRacklots(db *gorm.DB, racklotIDs []int) (bool, error) {
	if racklotIDs == nil {
		return false, nil
	}

	db = db.Where("racklot_id  in (?) ", racklotIDs).First(&AreaRacklot{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}
	return true, nil
}

func Create(db *gorm.DB, areaID int, racklotID int) (bool, error) {

	if areaID == 0 || racklotID == 0 {
		return false, errors.New("CreateAreaRacklot:AreaID or RacklotID  is Invalid")
	}

	if res, _ := IsExit(db, areaID, racklotID); res != false {
		return false, errors.New("CreateAreaRacklot Exist")
	}

	if err := db.Create(&AreaRacklot{RacklotID: racklotID, AreaID: areaID}).Error; err != nil {
		return false, err
	}

	return true, nil
}

func BatchCreate(db *gorm.DB, areaID int, racklots []int) (bool, error) {

	var data []interface{}
	for _, v := range racklots {
		var a AreaRacklot
		a.AreaID = areaID
		a.RacklotID = v

		data = append(data, a)
	}

	if err := gormbulk.BulkInsert(db, data, len(data)); err != nil {
		return false, err
	}

	return true, nil
}

func Update(db *gorm.DB, id int, updateMap map[string]interface{}) (bool, error) {

	db = db.Model(&AreaRacklot{}).Where("id = ?", id).Update(updateMap)

	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func Delete(db *gorm.DB, areaID int, racklotID int) (bool, error) {

	db = db.Where("area_id = ? and racklot_id = ?", areaID, racklotID).Unscoped().Delete(&AreaRacklot{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func DeleteByArea(db *gorm.DB, areaID int) (bool, error) {

	db = db.Where("area_id = ?", areaID).Unscoped().Delete(&AreaRacklot{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func DeleteByAreas(db *gorm.DB, areaIDs []int) (bool, error) {

	db = db.Where("area_id in (?)", areaIDs).Unscoped().Delete(&AreaRacklot{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func DeleteByRacklot(db *gorm.DB, racklotID int) (bool, error) {

	db = db.Where("racklot_id = ?", racklotID).Unscoped().Delete(&AreaRacklot{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func DeleteByRacklots(db *gorm.DB, racklotIDs []int) (bool, error) {

	db = db.Where("racklot_id in (?)", racklotIDs).Unscoped().Delete(&AreaRacklot{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func DeleteByAreaRacklots(db *gorm.DB, areaID int, racklotIDs []int) (bool, error) {

	db = db.Where("area_id = ? and racklot_id in (?)", areaID, racklotIDs).Unscoped().Delete(&AreaRacklot{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}
