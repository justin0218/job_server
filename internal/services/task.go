package services

import (
	"github.com/robfig/cron"
	"job_server/internal/models/task"
	"job_server/pkg/job"
)

func init() {
	TaskService := new(TaskService)
	TaskService.BillNotice(1)
}

type TaskService struct {
	baseService
}

//
func (s *TaskService) BillNotice(taskId int) {
	taskModel := task.NewModel(s.Mysql.Get())
	taskModel.InitLock(taskId)
	c := cron.New()
	spec := "0 0 18 * * ?"
	_ = c.AddFunc(spec, func() {
		err := taskModel.Lock(taskId)
		if err != nil {
			return
		}
		defer taskModel.UnLock(taskId)
		job.BillNotice()
	})
	c.Start()
}
