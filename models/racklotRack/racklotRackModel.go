package racklotRack

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// RacklotRackModel ..
type RacklotRack struct {
	gorm.Model
	RacklotID int
	RackID    int
}

func Query(db *gorm.DB, racklotID int, rackID int) (*RacklotRack, error) {

	if racklotID == 0 && rackID == 0 {
		return nil, nil
	}

	var m RacklotRack
	db = db.Where("racklot_id = ? or rack_id = ?", racklotID, rackID).First(&m)
	if err := db.Error; err != nil {
		return nil, err
	}

	if db.RowsAffected <= 0 {
		return nil, nil
	}

	return &m, nil
}

func IsExit(db *gorm.DB, racklotID int, rackID int) (bool, error) {
	if racklotID == 0 && rackID == 0 {
		return false, nil
	}

	db = db.Where("racklot_id = ? or rack_id = ?", racklotID, rackID).First(&RacklotRack{})
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

	db = db.Where("racklot_id  in (?) ", racklotIDs).First(&RacklotRack{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}
	return true, nil
}

func Create(db *gorm.DB, racklotID int, rackID int) (bool, error) {

	if racklotID == 0 && rackID == 0 {
		return false, nil
	}

	if r, _ := IsExit(db, racklotID, rackID); r != true {
		return false, errors.New("CreateRacklotRack Exist")
	}

	if err := db.Create(&RacklotRack{RackID: rackID, RacklotID: racklotID}).Error; err != nil {
		return false, err
	}

	return true, nil
}

func Update(db *gorm.DB, id int, updateMap map[string]interface{}) (bool, error) {

	db = db.Model(&RacklotRack{}).Where("id = ?", id).Update(updateMap)

	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func Delete(db *gorm.DB, racklotID int, rackID int) (bool, error) {

	db = db.Where("racklot_id = ? or rack_id = ?", racklotID, rackID).Unscoped().Delete(&RacklotRack{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func BatchDeleteByRacklots(db *gorm.DB, racklotIDs []int) (bool, error) {

	db = db.Where("racklot_id in (?) ", racklotIDs).Unscoped().Delete(&RacklotRack{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func BatchDeleteByRacks(db *gorm.DB, rackIDs []int) (bool, error) {

	db = db.Where("rack_id in (?) ", rackIDs).Unscoped().Delete(&RacklotRack{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}
