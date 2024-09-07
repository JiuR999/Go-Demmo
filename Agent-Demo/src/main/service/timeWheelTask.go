package service

import (
	"fmt"
	"github.com/rfyiamcool/go-timewheel"
	"time"
)

var timeWheel *timewheel.TimeWheel
var taskmap = map[string]*timewheel.Task{}

func InitTimeWheelOnStart() {
	StartTimeWheel()
	fmt.Println("StartTimeWheel finish")
}

func StartTimeWheel() {
	var err error
	timeWheel, err = timewheel.NewTimeWheel(500*time.Millisecond, 120, timewheel.TickSafeMode())
	if err != nil {
		panic(err)
	}
	timeWheel.Start()
}

func AddReportCronTask(id string, tick time.Duration, job func()) {
	fmt.Println("Add Cron job:", id)
	task := timewheel.AddCron(tick, job)
	taskmap[id] = task
}
