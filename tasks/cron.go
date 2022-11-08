package tasks

import (
	"giligili/util"
	"github.com/robfig/cron"
	"reflect"
	"runtime"
	"time"
)

var Cron *cron.Cron

// Run 运行
func Run(job func() error) {
	from := time.Now().UnixNano()
	err := job()
	to := time.Now().UnixNano()
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	if err != nil {
		util.Log().Error("%s fail：%dms", jobName, (to-from)/int64(time.Millisecond))
	} else {
		util.Log().Info("%s success：%dms", jobName, (to-from)/int64(time.Millisecond))
	}
}

// CronJob 定时任务
func CronJob() {
	if Cron == nil {
		Cron = cron.New()
	}
	if err := Cron.AddFunc("0 0 0 * * *", func() { Run(ClearDailyRank) }); err != nil {
		util.Log().Panic("定时任务%s启动失败，err：%s", "ClearDailyRank", err.Error())
	}
	Cron.Start()
	util.Log().Info("Cronjob start......")
}
