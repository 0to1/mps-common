package models

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

func QueryAreaRacklot(db *gorm.DB, AreaID int, RacklotID int) (*AreaRacklot, error) {

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

func CreateAreaRacklot(db *gorm.DB, areaID int, racklotID int) (bool, error) {

	if areaID == 0 && racklotID == 0 {
		return false, nil
	}

	if r, _ := QueryAreaRacklot(db, areaID, racklotID); r != nil {
		return false, errors.New("CreateAreaRacklot Exist")
	}

	if err := db.Create(&AreaRacklot{RacklotID: racklotID, AreaID: areaID}).Error; err != nil {
		return false, err
	}

	return true, nil
}

func BatchCreateAreaRacklot(db *gorm.DB, areaID int, racklots []int) (bool, error) {

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

func UpdateAreaRacklot(db *gorm.DB, id int, updateMap map[string]interface{}) (bool, error) {

	db = db.Model(&AreaRacklot{}).Where("id = ?", id).Update(updateMap)

	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func DeleteAreaRacklot(db *gorm.DB, areaID int, racklotID int) (bool, error) {

	db = db.Where("area_id = ? and racklot_id = ?", areaID, racklotID).Unscoped().Delete(&AreaRacklot{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func DeleteAreaRacklotByArea(db *gorm.DB, areaID int) (bool, error) {

	db = db.Where("area_id = ?", areaID).Unscoped().Delete(&AreaRacklot{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func DeleteAreaRacklotByRacklot(db *gorm.DB, areaID int, racklotID int) (bool, error) {

	db = db.Where("racklot_id = ?", racklotID).Unscoped().Delete(&AreaRacklot{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}
