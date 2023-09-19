package system

import (
	"dat320/lab4/scheduler/job"
	"time"
)

type Scheduler interface {
	Add(*job.Job)
	Tick(time.Duration) int
}
