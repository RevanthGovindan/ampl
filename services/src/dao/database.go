package dao

import (
	"errors"

	"gorm.io/gorm"
)

var DbConn DbPool

type DbPool struct {
	Db *gorm.DB
}

func (f *DbPool) GetAllTasks(t *[]Tasks, page int, limit int, count *int64) (err error) {
	var find *gorm.DB
	if limit != 0 && page != 0 {
		find = f.Db.Limit(limit).Offset(limit * (page - 1)).Find(t)
	} else {
		find = f.Db.Find(t)
	}
	err = find.Error
	err = errors.Join(err, f.Db.Model(&Tasks{}).Count(count).Error)
	if err != nil {
		return err
	}
	return nil
}

func (f *DbPool) SaveTask(t *Tasks) (err error) {
	txn := f.Db.Create(t)
	if txn.Error != nil {
		return nil
	}
	if txn.RowsAffected < 1 {
		return errors.New("insertion failed")
	}
	return nil
}

func (f *DbPool) GetTaskById(id string, t *Tasks) (err error) {
	var count int64
	find := f.Db.Find(t, id).Count(&count)
	if find.Error != nil {
		return err
	}
	if count == 0 {
		return errors.New("no data")
	}
	return nil
}

func (f *DbPool) UpdateTaskById(t Tasks) (err error) {
	find := f.Db.Model(&t).Updates(Tasks{Title: t.Title, Description: t.Description, Status: t.Status})
	if find.Error != nil {
		return err
	}
	if find.RowsAffected == 0 {
		return errors.New("invalid id or no changes")
	}
	return nil
}

func (f *DbPool) DeleteTaskById(id string) (err error) {
	find := f.Db.Delete(&Tasks{}, id)
	if find.Error != nil {
		return err
	}
	if find.RowsAffected == 0 {
		return errors.New("invalid id")
	}
	return nil
}
