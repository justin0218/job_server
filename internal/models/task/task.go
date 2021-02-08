package task

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Lock      int       `json:"lock"`
	Disabled  int       `json:"disabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Model struct {
	Db   *gorm.DB
	Name string
}

func NewModel(db *gorm.DB) *Model {
	return &Model{
		Db:   db,
		Name: "tasks",
	}
}

func (s *Model) InitLock(taskId int) {
	db := s.Db.Table(s.Name)
	db.Where("id = ?", taskId).Updates(map[string]interface{}{
		"lock": 1,
	})
	return
}

func (s *Model) Lock(taskId int) (err error) {
	db := s.Db.Table(s.Name)
	query := db.Where("id = ?", taskId)
	tk := new(Task)
	err = query.First(tk).Error
	if err != nil {
		return
	}
	if tk.Disabled != 0 {
		err = fmt.Errorf("lock is disabled！")
		return
	}
	err = query.UpdateColumn("lock", gorm.Expr("lock - ?", 1)).Error
	if err != nil {
		return
	}
	return
}

func (s *Model) UnLock(taskId int) (err error) {
	db := s.Db.Table(s.Name)
	query := db.Where("id = ?", taskId)
	tk := new(Task)
	err = query.First(tk).Error
	if err != nil {
		return
	}
	err = query.UpdateColumn("lock", 0).Error
	if err != nil {
		return
	}
	return
}
