package dao

import (
	"errors"

	"gorm.io/gorm"
)

var DbConn DbPool

type DbPool struct {
	Db *gorm.DB
}

func (f *DbPool) GetAllTasks(t *[]Tasks) (err error) {
	if err = f.Db.Find(t).Error; err != nil {
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
