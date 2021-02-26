package services

import (
	"github.com/robfig/cron"
	"job_server/internal/models/task"
	"job_server/pkg/job"
)

func init() {
	TaskService := new(TaskService)
	taskModel := task.NewModel(TaskService.Mysql.Get())

	TaskService.BillNotice(taskModel.InitLock(1))
	TaskService.AutoCloseOrder(taskModel.InitLock(2))

}

type TaskService struct {
	baseService
}

//
func (s *TaskService) BillNotice(taskLockModel *task.Model) {
	c := cron.New()
	spec := "0 0 18 * * ?"
	_ = c.AddFunc(spec, func() {
		err := taskLockModel.Lock()
		if err != nil {
			return
		}
		defer taskLockModel.UnLock()
		job.BillNotice()
	})
	c.Start()
}

func (s *TaskService) AutoCloseOrder(taskLockModel *task.Model) {
	c := cron.New()
	spec := "0/10 * * * * ?"
	_ = c.AddFunc(spec, func() {
		err := taskLockModel.Lock()
		if err != nil {
			return
		}
		defer taskLockModel.UnLock()
		job.AutoCloseOrder()
	})
	c.Start()
}
