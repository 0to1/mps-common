package cellMaterial

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// CellMaterialModel ..
type CellMaterial struct {
	gorm.Model
	CellID     uint32
	MaterialID int64
}

func Query(db *gorm.DB, cellID uint32, materialID int64) (*CellMaterial, error) {

	if cellID == 0 && materialID == 0 {
		return nil, nil
	}

	var m CellMaterial
	db = db.Where("cell_id = ? or material_id= ?", cellID, materialID).First(&m)
	if err := db.Error; err != nil {
		return nil, err
	}

	if db.RowsAffected <= 0 {
		return nil, nil
	}

	return &m, nil
}

func IsExit(db *gorm.DB, cellID uint32, materialID int64) (bool, error) {
	if cellID == 0 && materialID == 0 {
		return false, nil
	}

	db = db.Where("cell_id = ? or material_id= ?", cellID, materialID).First(&CellMaterial{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func IsExitByCells(db *gorm.DB, cellIDs []uint32) (bool, error) {
	if cellIDs == nil {
		return false, nil
	}

	db = db.Model(&CellMaterial{}).Where("cell_id  in (?) ", cellIDs).First(&CellMaterial{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}
	return true, nil
}

func Create(db *gorm.DB, cellID uint32, materialID int64) (bool, error) {

	if cellID == 0 || materialID == 0 {
		return false, errors.New("CreateCellMaterial:CellID or MaterialID  is Invalid")
	}

	if r, _ := IsExit(db, cellID, materialID); r != false {
		return false, errors.New("CreateCellMaterial Exist")
	}

	if err := db.Create(&CellMaterial{MaterialID: materialID, CellID: cellID}).Error; err != nil {
		return false, err
	}

	return true, nil
}

func Update(db *gorm.DB, id int, updateMap map[string]interface{}) (bool, error) {

	db = db.Model(&CellMaterial{}).Where("id = ?", id).Update(updateMap)

	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func Delete(db *gorm.DB, cellID uint32, materialID int64) (bool, error) {

	db = db.Where("cell_id = ? or material_id= ?", cellID, materialID).Unscoped().Delete(&CellMaterial{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func BatchDeleteByCells(db *gorm.DB, cellIDs []uint32) (bool, error) {

	db = db.Where("cell_id in (?) ", cellIDs).Unscoped().Delete(&CellMaterial{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func BatchDeleteByMaterials(db *gorm.DB, materialIDs []int64) (bool, error) {

	db = db.Where("material_id in (?) ", materialIDs).Unscoped().Delete(&CellMaterial{})
	if err := db.Error; err != nil {
		return false, err
	}

	if db.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}
