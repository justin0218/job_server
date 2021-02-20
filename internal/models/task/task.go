package task

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	LockFlag  int       `json:"lock_flag"`
	Disabled  int       `json:"disabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Model struct {
	Db     *gorm.DB
	Name   string
	TaskId int
}

func NewModel(db *gorm.DB) *Model {
	return &Model{
		Db:   db,
		Name: "tasks",
	}
}

func (s *Model) InitLock(taskId int) *Model {
	db := s.Db.Table(s.Name)
	db.Where("id = ?", taskId).Updates(map[string]interface{}{
		"lock_flag": 1,
	})
	s.TaskId = taskId
	return s
}

func (s *Model) Lock() (err error) {
	db := s.Db.Table(s.Name)
	query := db.Where("id = ?", s.TaskId)
	tk := new(Task)
	err = query.First(tk).Error
	if err != nil {
		return
	}
	if tk.Disabled != 0 {
		err = fmt.Errorf("lock is disabled！")
		return
	}
	err = query.UpdateColumn("lock_flag", gorm.Expr("lock_flag - ?", 1)).Error
	if err != nil {
		return
	}
	return
}

func (s *Model) UnLock() (err error) {
	db := s.Db.Table(s.Name)
	query := db.Where("id = ?", s.TaskId)
	tk := new(Task)
	err = query.First(tk).Error
	if err != nil {
		return
	}
	err = query.UpdateColumn("lock_flag", 1).Error
	if err != nil {
		return
	}
	return
}
