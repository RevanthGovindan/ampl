package service

import (
	"ampl/src/dao"
	"ampl/src/utils"
	"errors"

	"gorm.io/gorm"
)

type TaskService struct {
	Db *gorm.DB
}

func (f *TaskService) GetAllTasks(t *[]dao.Tasks, page int, limit int, count *int64) (err error) {
	var find *gorm.DB
	if limit != 0 && page != 0 {
		find = f.Db.Limit(limit).Offset(limit * (page - 1)).Find(t)
	} else {
		find = f.Db.Find(t)
	}
	err = find.Error
	err = errors.Join(err, f.Db.Model(&dao.Tasks{}).Count(count).Error)
	if err != nil {
		return err
	}
	return nil
}

func (f *TaskService) CreateTask(t *dao.Tasks) (err error) {
	txn := f.Db.Create(t)
	if txn.Error != nil {
		return nil
	}
	if txn.RowsAffected < 1 {
		return errors.New("insertion failed")
	}
	return nil
}

func (f *TaskService) GetTaskById(id uint64) (dao.Tasks, error) {
	var task dao.Tasks
	var count int64
	find := f.Db.Find(&task, id).Count(&count)
	if find.Error != nil {
		return task, find.Error
	}
	if count == 0 {
		return task, utils.NotFoundErr
	}
	return task, nil
}

func (f *TaskService) UpdateTaskById(t dao.Tasks) (dao.Tasks, error) {
	find := f.Db.Model(&t).Updates(dao.Tasks{Title: t.Title, Description: t.Description, Status: t.Status})
	if find.Error != nil {
		return t, find.Error
	}
	if find.RowsAffected == 0 {
		return t, errors.New("invalid id or no changes")
	}
	return t, nil
}

func (f *TaskService) DeleteTaskById(id uint64) (err error) {
	find := f.Db.Delete(&dao.Tasks{}, id)
	if find.Error != nil {
		return err
	}
	if find.RowsAffected == 0 {
		return errors.New("invalid id")
	}
	return nil
}
