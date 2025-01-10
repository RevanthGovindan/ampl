package orm

import (
	"ampl/src/config"
	"errors"
)

func GetAllTasks(t *[]Tasks) (err error) {
	if err = config.DbConn.Find(t).Error; err != nil {
		return err
	}
	return nil
}

func SaveTask(t *Tasks) (err error) {
	txn := config.DbConn.Create(t)
	if txn.Error != nil {
		return nil
	}
	if txn.RowsAffected < 1 {
		return errors.New("insertion failed")
	}
	return nil
}

func GetTaskById(id string, t *Tasks) (err error) {
	var count int64
	find := config.DbConn.Find(t, id).Count(&count)
	if find.Error != nil {
		return err
	}
	if count == 0 {
		return errors.New("no data")
	}
	return nil
}
