package task

import "github.com/jasonlvhit/gocron"

func PeriodicTasks(logCycle string) {

	gocron.Every(1).Day().Do(ClearLogData, logCycle)

	<-gocron.Start()
}
