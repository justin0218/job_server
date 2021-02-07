package job

import "github.com/robfig/cron"

func Run() {
	c := cron.New()
	spec := "0 0 18 * * ?"
	_ = c.AddFunc(spec, func() {
		BillNotice()
	})
	c.Start()
}
